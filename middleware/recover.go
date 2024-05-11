package middleware

import (
	"github.com/gin-gonic/gin"
	"social-todo-list/common"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.ErrorRes); ok {
					ctx.AbortWithError(appErr.StatusCode, appErr)
					return
				}

				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithError(appErr.StatusCode, appErr)
				return
			}
		}()

		ctx.Next()
	}
}
