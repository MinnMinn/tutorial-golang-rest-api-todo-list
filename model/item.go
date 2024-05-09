package todomodel

import (
	"errors"
	"social-todo-list/common"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type ToDoItem struct {
	common.SQLModel
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

func (item ToDoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}

	return nil
}
