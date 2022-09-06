package initx

import (
	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/pkg/errors"
	"github.com/cloudslit/newca/pkg/storage"
	"github.com/cloudslit/newca/pkg/storage/ipfs"
	web3_storage "github.com/cloudslit/newca/pkg/storage/web3-storage"
)

// InitStorage 初始化存储引擎
func InitStorage() (storage.ExecCloser, func(), error) {
	cfg := config.C
	var db storage.ExecCloser
	var err error
	switch cfg.Storage.Type {
	case "ipfs":
		db, err = ipfs.New()
		if err != nil {
			return nil, nil, err
		}
	case "web3.storage":
		db, err = web3_storage.New()
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("unknown storage")
	}
	return db, func() {
		if db != nil {
			db.Close()
		}
	}, nil
}
