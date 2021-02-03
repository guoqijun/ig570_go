package base

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"net/http"
)

func GetCompany(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	appG.Response(http.StatusOK, 0, nil)
	return
}
