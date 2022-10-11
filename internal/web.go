package internal

import (
	"net/http"

	"github.com/cloudslit/cfssl/ocsp"
	"github.com/cloudslit/deca/internal/api"
	"github.com/cloudslit/deca/internal/middleware"
	"github.com/gin-gonic/gin"

	"github.com/cloudslit/deca/internal/config"
	"github.com/cloudslit/deca/internal/router"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())
	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// Router register
	r.Register(app)

	return app
}

func InitOcspEngine(r *api.OcspAPI) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", ocsp.NewResponder(r, nil))
	return mux
}
