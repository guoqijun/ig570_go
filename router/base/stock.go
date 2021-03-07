package base

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"ig570_go/ig570"
	"net/http"
)

func GetCompany(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	resp, _ := ig570.GetCompanyList()
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyInfo(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")

	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}

	resp, _ := ig570.GetCompanyInfo(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}
