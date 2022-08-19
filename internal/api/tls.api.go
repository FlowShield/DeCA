package api

import (
	"github.com/cloudslit/cfssl/api"
	"github.com/cloudslit/cfssl/auth"
	"github.com/cloudslit/newca/internal/ginx"
	"github.com/cloudslit/newca/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var TlsSet = wire.NewSet(wire.Struct(new(TlsAPI), "*"))

type TlsAPI struct {
	TlsSrv *service.TlsSrv
}

// 利用cfssl内部handle
func (a *TlsAPI) Info(c *gin.Context) {
	a.TlsSrv.Info(c)
}

func (a *TlsAPI) AuthSign(c *gin.Context) {
	ctx := c.Request.Context()
	var params auth.AuthenticatedRequest
	if err := ginx.ParseJSON(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.TlsSrv.AuthSign(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	if err := api.SendResponse(c.Writer, result); err != nil {
		ginx.ResError(c, err)
		return
	}
}
