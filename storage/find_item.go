package todostorage

import (
	"context"
	"errors"
	todomodel "social-todo-list/model"

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

		return nil, err // other errors
	}

	return &itemData, nil
}