package api

import (
	"github.com/flowshield/deca/internal/ginx"
	"github.com/flowshield/deca/internal/service"
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

func (a *CertificateAPI) Revoke(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CertificateSrv.Revoke(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *CertificateAPI) Verify(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CertificateSrv.Verify(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}
