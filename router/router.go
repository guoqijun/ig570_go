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
	baseFunc.GET("/company-share", base.GetCompanyShare)

	baseFunc.GET("/company-additional", base.GetCompanyAdditional)
	baseFunc.GET("/company-finance-digest", base.GetCompanyFinanceDigest)
	baseFunc.GET("/company-profit", base.GetCompanyProfit)
	baseFunc.GET("/company-cash-flow", base.GetCompanyCashFlow)
	baseFunc.GET("/company-performance", base.GetCompanyPerformance)
	baseFunc.GET("/company-finance-index", base.GetCompanyFinanceIndex)
	baseFunc.GET("/company-qfq-factor", base.GetCompanyQfqFactor)
	baseFunc.GET("/company-hfq-factor", base.GetCompanyHfqFactor)

	baseFunc.GET("/real", base.GetTimeReal)
	baseFunc.GET("/real-level-5", base.GetTimeRealLevel5)
	baseFunc.GET("/real-one-by-one", base.GetTimeRealOneByOne)
	baseFunc.GET("/real-time-deal", base.GetTimeRealTimeDeal)
	baseFunc.GET("/real-big-deal", base.GetTimeRealBigDeal)

	return r
}
