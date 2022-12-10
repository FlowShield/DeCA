package initx

import (
	"context"
	"github.com/flowshield/deca/internal/config"
	"github.com/flowshield/deca/pkg/contract"
)

func InitEthClient(ctx context.Context) (*contract.EthClient, func(), error) {
	cfg := config.C.Contract
	client, err := contract.NewEthClient(ctx, &contract.Config{
		PrivateKey: cfg.PrivateKey,
		Address:    cfg.Address,
		RpcUrl:     cfg.RpcUrl,
	})
	if err != nil {
		return nil, nil, err
	}
	return client, func() {
		client.Client.Close()
	}, err
}
