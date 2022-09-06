package web3_storage

import (
	"context"
	"github.com/cloudslit/newca/internal/config"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"github.com/wumansgy/goEncrypt"
)

type Web3Storage struct {
	db  w3s.Client
	key []byte
}

func New() (*Web3Storage, error) {
	cfg := config.C.Web3Storage
	db, err := w3s.NewClient(w3s.WithToken(cfg.Token))
	if err != nil {
		return nil, err
	}
	return &Web3Storage{
		db:  db,
		key: []byte(cfg.EncryptKey),
	}, nil
}

func (a *Web3Storage) Put(ctx context.Context, data []byte) (cid string, err error) {
	file, err := a.dataToFile(data)
	defer os.Remove(file.Name())
	if err != nil {
		return
	}
	cidObj, err := a.db.Put(ctx, file)
	if err != nil {
		return
	}
	return cidObj.String(), nil
}

func (a *Web3Storage) Get(ctx context.Context, cidStr string) (data []byte, err error) {
	cidObj, err := cid.Decode(cidStr)
	if err != nil {
		return nil, err
	}
	res, err := a.db.Get(ctx, cidObj)
	if err != nil {
		return nil, err
	}
	_, fsys, err := res.Files()
	if err != nil {
		return nil, err
	}
	err = fs.WalkDir(fsys, "/", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			file, err := fsys.Open(path)
			if err != nil {
				return err
			}
			data, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	data, err = goEncrypt.DesCbcDecrypt(data, a.key, nil) //解密得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		return nil, err
	}
	return
}

func (a *Web3Storage) Close() error {
	return nil
}

func (a *Web3Storage) dataToFile(data []byte) (file *os.File, err error) {
	file, err = os.CreateTemp("", "data")
	if err != nil {
		return
	}
	// 对数据进行加密
	cryptText, err := goEncrypt.DesCbcEncrypt(data, a.key, nil)
	if err != nil {
		return
	}
	err = os.WriteFile(file.Name(), cryptText, 0644)
	if err != nil {
		return
	}
	return
}
