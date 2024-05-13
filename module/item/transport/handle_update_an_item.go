package todotrpt

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/item/business"
	"social-todo-list/module/item/model"
	"social-todo-list/module/item/storage"
	"strconv"
)

func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewUpdateToDoItemBiz(storage)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
