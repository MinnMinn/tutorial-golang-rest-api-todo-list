package todostorage

import (
	"context"
	"social-todo-list/common"
	todomodel "social-todo-list/model"
)

func (s *mysqlStorage) ListItem(
	ctx context.Context,
	filter *todomodel.Filter,
	paging *common.Paging,
) ([]todomodel.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []todomodel.ToDoItem

	if err := s.db.Table(todomodel.ToDoItem{}.TableName()).
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
