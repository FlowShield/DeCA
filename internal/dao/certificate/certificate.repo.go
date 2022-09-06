package certificate

import (
	"context"
	"github.com/IceFireDB/icefiredb-crdt-kv/kv"
	"github.com/cloudslit/newca/internal/schema"
	"github.com/cloudslit/newca/pkg/errors"
	"github.com/cloudslit/newca/pkg/storage"
	"github.com/cloudslit/newca/pkg/util/json"
	"github.com/google/wire"
)

var CertificateSet = wire.NewSet(wire.Struct(new(CertificateRepo), "*"))

type CertificateRepo struct {
	DB   storage.ExecCloser
	Crdt *kv.CRDTKeyValueDB
}

func (a *CertificateRepo) GetS(ctx context.Context, id string) (*schema.Certificate, error) {
	errCount := 0
retry:
	str, err := a.DB.Get(ctx, id)
	if err != nil {
		if errCount < 3 {
			errCount++
			goto retry
		}
		return nil, err
	}
	var result Certificate
	err = json.Unmarshal(str, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result.ToSchemaCertificate(), err
}

func (a *CertificateRepo) PutS(ctx context.Context, item schema.Certificate) (*schema.IDResult, error) {
	itemByte, err := json.Marshal(item)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	errCount := 0
retry:
	id, err := a.DB.Put(ctx, itemByte)
	if err != nil {
		if errCount < 3 {
			errCount++
			goto retry
		}
		return nil, errors.WithStack(err)
	}
	return schema.NewIDResult(id), nil
}

func (a *CertificateRepo) GetC(ctx context.Context, key string) ([]byte, error) {
	result, err := a.Crdt.Get(ctx, []byte(key))
	if err != nil {
		return nil, err
	}
	return result, err
}

func (a *CertificateRepo) PutC(ctx context.Context, key string, value string) error {
	return a.Crdt.Put(ctx, []byte(key), []byte(value))
}
