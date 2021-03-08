package base

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"ig570_go/ig570"
	"net/http"
)

func GetTimeHistoryTrade(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeHistoryTrade(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeHistoryKdj(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeHistoryKdj(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeHistoryMacd(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeHistoryMacd(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeHistoryMa(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeHistoryMa(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}

func GetTimeHistoryBoll(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	code := ctx.DefaultQuery("code", "")
	level := ctx.DefaultQuery("level", "")
	if len(code) == 0 || len(level) == 0 {
		appG.Response(http.StatusOK, -1, nil)
		return
	}
	resp, _ := ig570.GetTimeHistoryBoll(code, level)
	appG.Response(http.StatusOK, 0, resp)
	return
}
