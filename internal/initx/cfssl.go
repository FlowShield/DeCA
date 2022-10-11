package initx

import (
	"crypto"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cloudslit/cfssl/api/info"
	"github.com/cloudslit/cfssl/helpers"
	"github.com/cloudslit/cfssl/ocsp"
	"github.com/cloudslit/cfssl/signer"
	"github.com/cloudslit/cfssl/signer/local"
	"github.com/cloudslit/deca/internal/config"
)

// InitCert 初始化证书
func InitCert() (*x509.Certificate, crypto.Signer, error) {
	cfg := config.C
	certPEM, err := ioutil.ReadFile(cfg.TLS.CertFile)
	if err != nil {
		return nil, nil, err
	}
	keyPEM, err := ioutil.ReadFile(cfg.TLS.KeyFile)
	if err != nil {
		return nil, nil, err
	}
	key, err := helpers.ParsePrivateKeyPEM(keyPEM)
	if err != nil {
		return nil, nil, err
	}
	cert, err := helpers.ParseCertificatePEM(certPEM)
	if err != nil {
		return nil, nil, err
	}
	return cert, key, nil
}

type CfsslHandler struct {
	LocalSigner   signer.Signer
	OcspSigner    ocsp.Signer
	InfoHandler   http.Handler
	RevokeHandler http.Handler
}

func InitCfssl() (*CfsslHandler, error) {
	cfg := config.C
	cert, key, err := InitCert()
	if err != nil {
		return nil, err
	}

	// localSigner
	localSigner, err := local.NewSigner(key, cert, signer.DefaultSigAlgo(key), cfg.Cfssl.Config.Signing)
	if err != nil {
		return nil, err
	}

	// ocspSigner
	ocspSigner, err := ocsp.NewSigner(cert, cert, key, 4*24*time.Hour)
	if err != nil {
		return nil, err
	}

	// infoHandler
	infoHandler, err := info.NewHandler(localSigner)
	if err != nil {
		return nil, err
	}

	return &CfsslHandler{
		LocalSigner: localSigner,
		OcspSigner:  ocspSigner,
		InfoHandler: infoHandler,
	}, nil
}
