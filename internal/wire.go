//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package internal

import (
	"context"
	"github.com/flowshield/deca/internal/service"
	"github.com/google/wire"

	"github.com/flowshield/deca/internal/api"
	"github.com/flowshield/deca/internal/dao"
	"github.com/flowshield/deca/internal/initx"
	"github.com/flowshield/deca/internal/router"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		initx.InitCfssl,
		initx.InitStorage,
		initx.InitEthClient,
		//initx.InitCrdtKv,
		initx.InitOcspCache,
		dao.RepoSet,
		service.ServiceSet,
		InitGinEngine,
		InitOcspEngine,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
