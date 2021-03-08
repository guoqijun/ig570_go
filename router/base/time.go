package base

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"ig570_go/ig570"
	"net/http"
)

func GetTimeRealLevel(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealLevel(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealKdj(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealKdj(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealMacd(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealMacd(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealMa(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealMa(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeRealBoll(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeRealBoll(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}
