package repository

import (
	"context"

	"github.com/ag9920/learnfx/domain/entity"
)

type ItemRepo interface {
	CreateItem(ctx context.Context, item *entity.Item) (int64, error)
	BatchQueryItemByID(ctx context.Context, itemIDs []int64) ([]*entity.Item, error)
}
