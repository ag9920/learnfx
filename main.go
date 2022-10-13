package main

import (
	"go.uber.org/fx"

	"github.com/ag9920/learnfx/domain"
	"github.com/ag9920/learnfx/infra"
)

func main() {
	fx.New(
		domain.Module,
		infra.Module,
		fx.Provide(NewRpcServer, NewLearnFxServiceImpl),
		fx.Invoke(startServer),
	).Run()
}
