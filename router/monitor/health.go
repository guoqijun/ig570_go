package Health

import (
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"net/http"
)

type Monitor struct {
}

type Health interface {
	Health(*gin.Context)
}

func NewHealth() Health {
	return &Monitor{}
}

func (ret *Monitor) Health(ctx *gin.Context) {
	appG := common.Gin{C: ctx}
	appG.Response(http.StatusOK, 0, nil)
}
