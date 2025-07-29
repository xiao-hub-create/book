package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": err.Error()})
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}
