package router

import (
	"github.com/gin-gonic/gin"
	"ig570_go/conf"
	"ig570_go/router/base"
	Health "ig570_go/router/monitor"
)

func RoutersInit() *gin.Engine {
	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()

	MonitorHealth := Health.NewMonitor()

	monitorRoute := r.Group("/health")
	monitorRoute.GET("/", MonitorHealth.Health)

	baseFunc := r.Group("/base")
	baseFunc.GET("/company", base.GetCompany)
	baseFunc.GET("/company-info", base.GetCompanyInfo)

	return r
}
