package initx

import (
	"context"
	"github.com/IceFireDB/icefiredb-crdt-kv/kv"
	"github.com/cloudslit/newca/internal/config"
	"github.com/cloudslit/newca/pkg/logger"
)

func InitCrdtKv(ctx context.Context) (*kv.CRDTKeyValueDB, func(), error) {
	cfg := config.C.CrdtKv
	log := logger.StandardLogger()
	db, err := kv.NewCRDTKeyValueDB(ctx, kv.Config{
		NodeServiceName:     cfg.NodeServiceName,
		DataStorePath:       cfg.DataStorePath,
		DataSyncChannel:     cfg.DataSyncChannel,
		NetDiscoveryChannel: cfg.NetDiscoveryChannel,
		Namespace:           cfg.Namespace,
		Logger:              log,
	})
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		db.Close()
	}, err
}
