package service

import (
	"context"

	"github.com/cloudslit/deca/internal/dao"
	"github.com/cloudslit/deca/internal/schema"
	"github.com/cloudslit/deca/pkg/errors"
	"github.com/google/wire"
)

var CertificateSet = wire.NewSet(wire.Struct(new(CertificateSrv), "*"))

type CertificateSrv struct {
	CertificateRepo *dao.CertificateRepo
}

func (a *CertificateSrv) GetS(ctx context.Context, id string) (*schema.Certificate, error) {
	item, err := a.CertificateRepo.GetS(ctx, id)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *CertificateSrv) GetC(ctx context.Context, key string) (interface{}, error) {
	item, err := a.CertificateRepo.GetC(ctx, schema.SnCidKey(key))
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return string(item), nil
}
