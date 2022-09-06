package internal

import (
	"github.com/cloudslit/cfssl/ocsp"
	"github.com/cloudslit/newca/internal/api"
	"github.com/cloudslit/newca/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/internal/router"
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
