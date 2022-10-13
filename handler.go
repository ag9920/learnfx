package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/server"
	"go.uber.org/fx"

	"github.com/ag9920/learnfx/domain/entity"
	"github.com/ag9920/learnfx/domain/service"
	"github.com/ag9920/learnfx/kitex_gen/learn/fx/item"
	"github.com/ag9920/learnfx/kitex_gen/learn/fx/item/learnfxservice"
)

func NewRpcServer(rpcSrv item.LearnFxService) server.Server {
	return learnfxservice.NewServer(rpcSrv)
}

func startServer(svr server.Server, lc fx.Lifecycle) {
	// 通过lifecycle异步启动，不然invoke执行在onstart之前
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("start rpc server")
			// 异步启动避免阻塞fx启动，依赖srv.Run的panic
			go func() {
				if err := svr.Run(); err != nil {
					fmt.Printf("fail to run server: %v", err)
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("shutdown rpc server")
			return svr.Stop()
		},
	})
}

type LearnFxServiceImpl struct {
	ItemDomainService service.ItemDomainService
}

func NewLearnFxServiceImpl(s service.ItemDomainService) item.LearnFxService {
	return &LearnFxServiceImpl{
		ItemDomainService: s,
	}
}

func (i *LearnFxServiceImpl) FilterVisibleItems(ctx context.Context, req *item.FilterVisibleItemsReq) (*item.FilterVisibleItemsResp, error) {
	visibleItemIDs, err := i.ItemDomainService.FilterVisibleItems(ctx, req.ItemIDs, req.UserID)
	if err != nil {
		return nil, err
	}
	resp := &item.FilterVisibleItemsResp{
		VisibleItemIDs: visibleItemIDs,
	}
	return resp, nil
}

func (i *LearnFxServiceImpl) CreateItem(ctx context.Context, req *item.CreateItemReq) (resp *item.CreateItemResp, err error) {
	itemID, err := i.ItemDomainService.CreateItem(ctx, &entity.Item{
		ID:           req.Item.ID,
		Name:         req.Item.Name,
		Desc:         req.Item.Desc,
		VisibleUsers: req.Item.VisibleUsers,
	})
	if err != nil {
		return nil, err
	}
	resp = &item.CreateItemResp{
		ID: itemID,
	}
	return resp, nil
}
