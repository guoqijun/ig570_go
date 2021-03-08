package ig570

import (
	"ig570_go/common"
	"ig570_go/conf"
)

/**
股票列表接口
*/
func GetCompanyList() (response string, err error) {
	requestUrl := "http://ig507.com/data/base/gplist"
	return common.Request(requestUrl)
}

/**
公司简介
*/
func GetCompanyInfo(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/info/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
近年分红
*/
func GetCompanyShare(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/share/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
近年增发
*/
func GetCompanyAdditional(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/zf/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
财务摘要
*/
func GetCompanyFinanceDigest(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/fs/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
近一年各季度利润
*/
func GetCompanyProfit(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/pf/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
近一年各季度现金流
*/
func GetCompanyCashFlow(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/cf/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
近年业绩预告
*/
func GetCompanyPerformance(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/ep/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
财务指标
*/
func GetCompanyFinanceIndex(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/f10/fi/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
日K线前复权因子
*/
func GetCompanyQfqFactor(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/qfq_factor/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
日K线后复权因子
*/
func GetCompanyHfqFactor(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/hfq_factor/ " + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}
