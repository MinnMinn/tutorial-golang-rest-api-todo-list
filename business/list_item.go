package todobiz

import (
	"context"
	"social-todo-list/common"
	todomodel "social-todo-list/model"
)

type ListTodoItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *todomodel.Filter,
		paging *common.Paging,
	) ([]todomodel.ToDoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListToDoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItems(ctx context.Context,
	filter *todomodel.Filter,
	paging *common.Paging,
) ([]todomodel.ToDoItem, error) {
	result, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
