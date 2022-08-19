//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"context"
	"github.com/cloudslit/newca/internal/service"
	"github.com/google/wire"

	"github.com/cloudslit/newca/internal/api"
	"github.com/cloudslit/newca/internal/dao"
	"github.com/cloudslit/newca/internal/initx"
	"github.com/cloudslit/newca/internal/router"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		initx.InitSigner,
		initx.InitInfoHandle,
		initx.InitStorage,
		initx.InitCrdtKv,
		dao.RepoSet,
		service.ServiceSet,
		InitGinEngine,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
