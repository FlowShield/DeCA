package service

import (
	"context"
	"encoding/hex"
	"github.com/cloudslit/cfssl/helpers"
	"github.com/cloudslit/cfssl/ocsp"
	"github.com/cloudslit/newca/internal/dao"
	"github.com/cloudslit/newca/internal/initx"
	"github.com/cloudslit/newca/internal/schema"
	"github.com/cloudslit/newca/pkg/errors"
	"github.com/cloudslit/newca/pkg/logger"
	"github.com/cloudslit/newca/pkg/memorycacher"
	"github.com/cloudslit/newca/pkg/util/json"

	"github.com/google/wire"
	stdocsp "golang.org/x/crypto/ocsp"
	"net/http"
)

var OcspSet = wire.NewSet(wire.Struct(new(OcspSrv), "*"))

type OcspSrv struct {
	CertificateRepo *dao.CertificateRepo
	CfsslHandler    *initx.CfsslHandler
	Cache           *memorycacher.Cache
	Ctx             context.Context
}

func (a *OcspSrv) Query(req *stdocsp.Request) ([]byte, http.Header, error) {
	if req == nil {
		return nil, nil, errors.New("called with nil request")
	}
	aki := hex.EncodeToString(req.IssuerKeyHash)
	sn := req.SerialNumber

	if sn == nil {
		return nil, nil, errors.New("request contains no serial")
	}
	strSN := sn.String()

	if cachedResp, ok := a.Cache.Get(strSN + aki); ok {
		if resp, ok := cachedResp.([]byte); ok {
			return resp, nil, nil
		}
		logger.Errorf("cache值解析错误, sn:%s, aki:%s", strSN, aki)
	}

	// 数据库查询
	crdtId, err := a.CertificateRepo.GetC(a.Ctx, schema.SnCidKey(strSN))
	if err != nil {
		return nil, nil, err
	}
	certS, err := a.CertificateRepo.GetS(a.Ctx, string(crdtId))
	if err != nil {
		return nil, nil, err
	}

	cert, err := helpers.ParseCertificatePEM([]byte(certS.CertPem))
	if err != nil {
		logger.WithErrorStack(a.Ctx, err).Errorf("证书PEM解析错误:%s, sn:%s, aki:%s", err.Error(), strSN, aki)
		return nil, nil, err
	}
	signReq := &ocsp.SignRequest{
		Certificate: cert,
		Status:      "good",
		Reason:      0,
	}
	// 查询是否吊销
	revoke, _ := a.CertificateRepo.GetC(a.Ctx, schema.SnRevokeKey(strSN))
	if revoke != nil {
		var re schema.CertificateRevoke
		_ = json.Unmarshal(revoke, &re)
		signReq.Status = "revoked"
		signReq.Reason = 1
		signReq.RevokedAt = re.RevokeAt
	}
	ocspResp, err := a.CfsslHandler.OcspSigner.Sign(*signReq)
	if err != nil {
		logger.WithErrorStack(a.Ctx, err).Errorf("OCSP签名错误:%s, sn:%s, aki:%s", err.Error(), strSN, aki)
		return nil, nil, err
	}
	a.Cache.SetDefault(strSN+aki, ocspResp)

	logger.Infof("OCSP签名完成, sn:%s, aki:%s", strSN, aki)
	return ocspResp, nil, nil
}
