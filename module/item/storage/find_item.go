package todostorage

import (
	"context"
	"errors"
	"social-todo-list/common"
	"social-todo-list/module/item/model"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(
	ctx context.Context,
	condition map[string]interface{},
) (*todomodel.ToDoItem, error) {
	var itemData todomodel.ToDoItem

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // data not found
			return nil, todomodel.ErrItemNotFound
		}

		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &itemData, nil
}
