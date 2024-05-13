package todostorage

import (
	"context"
	"social-todo-list/common"
	todomodel2 "social-todo-list/module/item/model"
)

func (s *mysqlStorage) ListItem(
	ctx context.Context,
	filter *todomodel2.Filter,
	paging *common.Paging,
) ([]todomodel2.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []todomodel2.ToDoItem

	if err := s.db.Table(todomodel2.ToDoItem{}.TableName()).
		Where(&filter).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
