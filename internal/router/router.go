package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/cloudslit/newca/internal/api"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	TlsAPI         *api.TlsAPI
	CertificateAPI *api.CertificateAPI
} // end

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		gCfssl := v1.Group("cfssl")
		{
			gCfssl.POST("info", a.TlsAPI.Info)
			gCfssl.POST("authsign", a.TlsAPI.AuthSign)
		}

		gCert := v1.Group("certificate")
		{
			gCert.GET("gets/:id", a.CertificateAPI.GetS)
			gCert.GET("getc/:id", a.CertificateAPI.GetC)
		}
	} // v1 end
}
