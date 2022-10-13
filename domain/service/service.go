package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ag9920/learnfx/domain/entity"
	"github.com/ag9920/learnfx/domain/repository"
)

type ItemDomainService interface {
	CreateItem(ctx context.Context, item *entity.Item) (int64, error)
	FilterVisibleItems(ctx context.Context, itemIDs []int64, userID int64) ([]int64, error)
}

type ItemDomainServiceImpl struct {
	ItemRepo repository.ItemRepo
}

var _ ItemDomainService = new(ItemDomainServiceImpl)

func NewItemDomainServiceImpl(repo repository.ItemRepo) ItemDomainService {
	return &ItemDomainServiceImpl{
		ItemRepo: repo,
	}
}

func (i *ItemDomainServiceImpl) CreateItem(ctx context.Context, item *entity.Item) (int64, error) {
	if item.ID == 0 {
		item.ID = time.Now().Unix()
	}
	return i.ItemRepo.CreateItem(ctx, item)
}

func (i *ItemDomainServiceImpl) FilterVisibleItems(ctx context.Context, itemIDs []int64, userID int64) ([]int64, error) {
	items, err := i.ItemRepo.BatchQueryItemByID(ctx, itemIDs)
	if err != nil {
		return nil, fmt.Errorf("ItemRepo.BatchQueryItemByID failed, err=%w", err)
	}

	result := make([]int64, 0, len(items))
	for _, item := range items {
		var isVisible bool
		for _, visibleUserID := range item.VisibleUsers {
			if userID == visibleUserID {
				isVisible = true
				break
			}
		}
		if isVisible {
			result = append(result, item.ID)
		}
	}
	return result, nil
}
