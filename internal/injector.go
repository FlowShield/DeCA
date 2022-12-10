package internal

import (
	"github.com/flowshield/deca/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine     *gin.Engine
	OcspEngine *http.ServeMux
	Router     *router.Router
}
