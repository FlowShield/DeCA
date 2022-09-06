package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/pkg/errors"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

type Ipfs struct {
	sh *shell.Shell
}

func New() (*Ipfs, error) {
	cfg := config.C.Ipfs
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	sh := shell.NewShell(addr)
	if !sh.IsUp() {
		return nil, errors.New("ipfs is not up")
	}
	return &Ipfs{
		sh: sh,
	}, nil
}

func (a *Ipfs) Put(ctx context.Context, data []byte) (string, error) {
	hash, err := a.sh.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (a *Ipfs) Get(ctx context.Context, hash string) ([]byte, error) {
	read, err := a.sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(read)
	return body, nil
}

func (a *Ipfs) Close() error {
	return nil
}
