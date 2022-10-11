//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package internal

import (
	"context"
	"github.com/cloudslit/deca/internal/service"
	"github.com/google/wire"

	"github.com/cloudslit/deca/internal/api"
	"github.com/cloudslit/deca/internal/dao"
	"github.com/cloudslit/deca/internal/initx"
	"github.com/cloudslit/deca/internal/router"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		initx.InitCfssl,
		initx.InitStorage,
		initx.InitCrdtKv,
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
