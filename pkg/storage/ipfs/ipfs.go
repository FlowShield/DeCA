package ipfs

import (
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

type Ipfs struct {
	sh *shell.Shell
}

func New(sh *shell.Shell) *Ipfs {
	return &Ipfs{
		sh: sh,
	}
}

func (a *Ipfs) Put(data string) (string, error) {
	hash, err := a.sh.Add(bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (a *Ipfs) Get(hash string) (string, error) {
	read, err := a.sh.Cat(hash)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)
	return string(body), nil
}

func (a *Ipfs) Close() error {
	return nil
}
