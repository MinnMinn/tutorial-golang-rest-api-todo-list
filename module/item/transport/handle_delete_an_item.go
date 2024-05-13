package todotrpt

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/item/business"
	"social-todo-list/module/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleDeleteAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewDeleteToDoItemBiz(storage)

		if err := biz.DeleteItem(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
