package initx

import (
	"github.com/cloudslit/cfssl/api/info"
	cfssl_config "github.com/cloudslit/cfssl/config"
	"github.com/cloudslit/cfssl/helpers"
	"github.com/cloudslit/cfssl/signer"
	"github.com/cloudslit/cfssl/signer/local"
	"github.com/cloudslit/newca/internal/config"
	"io/ioutil"
	"net/http"
)

// InitSigner 初始化cfssl Signer
func InitSigner() (signer.Signer, error) {
	cfg := config.C
	certPEM, err := ioutil.ReadFile(cfg.App.CertFile)
	if err != nil {
		return nil, err
	}
	keyPEM, err := ioutil.ReadFile(cfg.App.KeyFile)
	if err != nil {
		return nil, err
	}
	key, err := helpers.ParsePrivateKeyPEM(keyPEM)
	if err != nil {
		return nil, err
	}
	cert, err := helpers.ParseCertificatePEM(certPEM)
	if err != nil {
		return nil, err
	}
	cfcfg, err := cfssl_config.LoadFile(cfg.Cfssl.ConfigFile)
	if err != nil {
		return nil, err
	}
	return local.NewSigner(key, cert, signer.DefaultSigAlgo(key), cfcfg.Signing)
}

// InitInfoHandle 初始化cfssl info Handle
func InitInfoHandle(s signer.Signer) (http.Handler, error) {
	return info.NewHandler(s)
}
