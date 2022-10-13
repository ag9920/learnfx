package domain

import (
	"go.uber.org/fx"

	"github.com/ag9920/learnfx/domain/service"
)

var Module = fx.Module("domain",
	fx.Provide(service.NewItemDomainServiceImpl),
)
