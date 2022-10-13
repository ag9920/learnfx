package infra

import (
	"context"

	"github.com/ag9920/learnfx/domain/entity"
	"github.com/ag9920/learnfx/domain/repository"
)

type ItemRepoImpl struct {
	ItemStore ItemMapStore
}

var _ repository.ItemRepo = new(ItemRepoImpl)

func NewItemRepoImpl(store ItemMapStore) repository.ItemRepo {
	return &ItemRepoImpl{
		ItemStore: store,
	}
}

func (i *ItemRepoImpl) CreateItem(ctx context.Context, item *entity.Item) (int64, error) {
	return i.ItemStore.StoreItem(item)
}

func (i *ItemRepoImpl) BatchQueryItemByID(ctx context.Context, itemIDs []int64) ([]*entity.Item, error) {
	result := make([]*entity.Item, 0, len(itemIDs))
	for _, itemID := range itemIDs {
		item := i.ItemStore.FindItemByID(itemID)
		if item != nil {
			result = append(result, item)
		}
	}
	return result, nil
}
