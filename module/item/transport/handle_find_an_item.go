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

func HandleFindAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewFindToDoItemBiz(storage)

		data, err := biz.FindAnItem(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
