package plugin

import (
	"github.com/nori-io/nori-common/v2/config"
	"github.com/nori-io/nori-common/v2/logger"
	"github.com/nori-io/nori-common/v2/plugin"
	"github.com/nori-io/nori-common/v2/storage"
	"github.com/nori-io/nori/internal/domain/manager"
	"github.com/nori-io/nori/internal/domain/repository"
	"go.uber.org/fx"
)

type ManagerParams struct {
	fx.In

	PluginRepository repository.PluginRepository
	RegistryService  plugin.Registry
	ConfigManager    config.Manager
	Storage          storage.Storage
	Logger           logger.Logger
}

func New(params ManagerParams) (manager.Plugin, error) {
	bucket, err := params.Storage.CreateBucketIfNotExists("plugins")
	if err != nil {
		return nil, err
	}

	return &Manager{
		pluginRepository: params.PluginRepository,
		registryService:  params.RegistryService,
		cm:               params.ConfigManager,
		logger:           params.Logger,
		bucket:           bucket,
	}, nil
}
