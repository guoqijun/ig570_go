package base

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"ig570_go/ig570"
	"net/http"
)

func GetTimeReal(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeReal(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealLevel5(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealLevel5(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealOneByOne(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealOneByOne(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealTimeDeal(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealTimeDeal(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealBigDeal(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	if len(code) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealBigDeal(code)
	appG.Response(http.StatusOK, 0, resp)
	return
}
