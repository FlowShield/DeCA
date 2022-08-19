package api

import (
	"github.com/cloudslit/newca/internal/ginx"
	"github.com/cloudslit/newca/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CertificateSet = wire.NewSet(wire.Struct(new(CertificateAPI), "*"))

type CertificateAPI struct {
	CertificateSrv *service.CertificateSrv
}

func (a *CertificateAPI) GetS(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CertificateSrv.GetS(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *CertificateAPI) GetC(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CertificateSrv.GetC(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}
