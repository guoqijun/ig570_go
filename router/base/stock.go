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

func GetCompanyShare(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyShare(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyAdditional(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyAdditional(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyFinanceDigest(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyFinanceDigest(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyProfit(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyProfit(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyCashFlow(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyCashFlow(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyPerformance(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyPerformance(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetCompanyFinanceIndex(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetCompanyFinanceIndex(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}
