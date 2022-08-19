package initx

import (
	"fmt"
	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/pkg/errors"
	"github.com/cloudslit/newca/pkg/storage"
	"github.com/cloudslit/newca/pkg/storage/ipfs"
	shell "github.com/ipfs/go-ipfs-api"
)

// InitStorage 初始化存储引擎
func InitStorage() (storage.ExecCloser, func(), error) {
	cfg := config.C
	var db storage.ExecCloser
	cleanFunc := func() {}
	switch cfg.Storage.Type {
	case "ipfs":
		addr := fmt.Sprintf("%s:%d", cfg.Ipfs.Host, cfg.Ipfs.Port)
		sh := shell.NewShell(addr)
		if !sh.IsUp() {
			return nil, nil, errors.New("ipfs is not up")
		}
		db = ipfs.New(sh)
	default:
		return nil, nil, errors.New("unknown storage")
	}

	return db, cleanFunc, nil
}
