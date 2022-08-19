package service

import (
	"context"
	"encoding/hex"
	"github.com/cloudslit/cfssl/auth"
	"github.com/cloudslit/cfssl/config"
	"github.com/cloudslit/cfssl/helpers"
	"github.com/cloudslit/cfssl/signer"
	"github.com/cloudslit/newca/internal/dao"
	"github.com/cloudslit/newca/internal/schema"
	"github.com/cloudslit/newca/pkg/attrmgr"
	"github.com/cloudslit/newca/pkg/errors"
	"github.com/cloudslit/newca/pkg/logger"
	"github.com/cloudslit/newca/pkg/util/json"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"time"
)

var TlsSet = wire.NewSet(wire.Struct(new(TlsSrv), "*"))

type TlsSrv struct {
	CertificateRepo *dao.CertificateRepo
	Signer          signer.Signer
	InfoHandle      http.Handler
}

func (a *TlsSrv) Info(c *gin.Context) {
	a.InfoHandle.ServeHTTP(c.Writer, c.Request)
}

func (a *TlsSrv) AuthSign(ctx context.Context, params auth.AuthenticatedRequest) (*schema.TlsShowResult, error) {
	var lsr schema.LocalSignRequest
	err := json.Unmarshal(params.Request, &lsr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Sanity checks to ensure that we have a valid policy. This
	// should have been checked in NewAuthHandler.
	policy := a.Signer.Policy()
	if policy == nil {
		return nil, errors.New("invalid policy")
	}

	profile, err := signer.Profile(a.Signer, lsr.Profile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if profile.Provider == nil {
		return nil, errors.New("no authentication provider")
	}

	validAuth := false
	if profile.Provider.Verify(&params) {
		validAuth = true
	} else if profile.PrevProvider != nil && profile.PrevProvider.Verify(&params) {
		validAuth = true
	}
	if !validAuth {
		return nil, errors.New("invalid token")
	}

	signReq := lsr.ToSignerSignRequest()
	if signReq.Request == "" {
		return nil, errors.New("missing parameter 'certificate_rest'")
	}
	if v, ok := signReq.Metadata["unique_id"]; ok {
		if v == "" {
			return nil, errors.New("Metadata unique_id required")
		}
	} else {
		return nil, errors.New("Metadata unique_id required")
	}
	err = a.genExpiryByCsr(&signReq, profile)
	if err != nil {
		return nil, err
	}
	cert, err := a.Signer.Sign(signReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 存储
	// save storage
	item, err := a.saveCertificate(ctx, cert, signReq.Metadata)
	if err != nil {
		return nil, err
	}
	logger.Infof("cid:%s", item.ID)
	return &schema.TlsShowResult{
		ID:          item.ID,
		Certificate: string(cert),
	}, nil
}

// 存储
func (a *TlsSrv) saveCertificate(ctx context.Context, certPem []byte, metaData map[string]interface{}) (*schema.IDResult, error) {
	cert, err := helpers.ParseCertificatePEM(certPem)
	if err != nil {
		return nil, err
	}
	certificate := schema.Certificate{
		SerialNumber:           cert.SerialNumber.String(),
		AuthorityKeyIdentifier: hex.EncodeToString(cert.AuthorityKeyId),
		CertPem:                string(certPem),
		NotBefore:              cert.NotBefore,
		NotAfter:               cert.NotAfter,
		MetaData:               metaData,
	}
	idResult, err := a.CertificateRepo.PutS(ctx, certificate)
	if err != nil {
		return nil, err
	}
	err = a.CertificateRepo.PutC(ctx, metaData["unique_id"].(string), idResult.ID)
	if err != nil {
		return nil, err
	}
	return schema.NewIDResult(idResult.ID), nil
}

// 添加过期时间
func (a *TlsSrv) genExpiryByCsr(sr *signer.SignRequest, profile *config.SigningProfile) error {
	csr, err := helpers.ParseCSRPEM([]byte(sr.Request))
	if err != nil {
		return errors.WithStack(err)
	}
	if v := attrmgr.GetExpiryValue(csr); v > 0 {
		var backdate time.Duration
		if backdate = profile.Backdate; backdate == 0 {
			backdate = -5 * time.Minute
		} else {
			backdate = -1 * profile.Backdate
		}
		notBefore := time.Now().Round(time.Minute).Add(backdate)
		notBefore = notBefore.UTC()
		notAfter := notBefore.Add(v)
		notAfter = notAfter.UTC()
		sr.NotBefore = notBefore
		sr.NotAfter = notAfter
	}
	return nil
}
