package permissions

import (
	"github.com/jamesnyakush/frika/internal/config"
	"github.com/jamesnyakush/frika/internal/store"
)

type Resource struct {
	Store  store.Store
	Config config.Config
}

func NewResource(store store.Store, cfg config.Config) *Resource {
	return &Resource{
		Store:  store,
		Config: cfg,
	}
}
