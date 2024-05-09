package todotrpt

import (
	"net/http"
	todobiz "social-todo-list/business"
	"social-todo-list/common"
	todomodel "social-todo-list/model"
	todostorage "social-todo-list/storage"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)

		// setup dependencies
		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewCreateToDoItemBiz(storage)

		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataItem.Id))
	}
}