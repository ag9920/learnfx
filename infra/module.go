package infra

import "go.uber.org/fx"

var Module = fx.Module("infra",
	fx.Provide(NewItemMapStore),
	fx.Provide(NewItemRepoImpl),
)
