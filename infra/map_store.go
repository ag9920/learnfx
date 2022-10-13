package infra

import (
	"github.com/ag9920/learnfx/domain/entity"
)

type ItemMapStore map[int64]*entity.Item

func NewItemMapStore() ItemMapStore {
	return make(map[int64]*entity.Item)
}

func (s ItemMapStore) StoreItem(item *entity.Item) (id int64, err error) {
	s[item.ID] = item
	return item.ID, nil
}

func (s ItemMapStore) FindItemByID(id int64) *entity.Item {
	return s[id]
}
