package Health

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"net/http"
)

type Monitor struct {
}

type IMonitor interface {
	Health(*gin.Context)
}

func NewMonitor() IMonitor {
	return &Monitor{}
}

func (ret *Monitor) Health(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	appG.Response(http.StatusOK, 0, "i am health")
}
