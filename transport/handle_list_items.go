package todotrpt

import (
	"net/http"
	todobiz "social-todo-list/business"
	"social-todo-list/common"
	todomodel "social-todo-list/model"
	todostorage "social-todo-list/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleListItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			todomodel.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		queryString.Process()

		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewListToDoItemBiz(storage)

		result, err := biz.ListItems(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, queryString.Paging, &queryString.Filter))
	}
}
