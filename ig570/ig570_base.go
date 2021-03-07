package ig570

import (
	"fmt"
	"ig570_go/conf"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
股票列表接口
*/
func GetCompanyList() (response string, err error) {
	getCompanyUrl := "http://ig507.com/data/base/gplist"
	return commonRequest(getCompanyUrl)
}

/**
公司简介
*/
func GetCompanyInfo(code string) (response string, err error) {
	getCompanyInfo := "http://ig507.com/data/time/f10/info/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyInfo)
}

/**
近年分红
*/
func GetCompanyShare(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/share/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
近年增发
*/
func GetCompanyAdditional(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/zf/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
财务摘要
*/
func GetCompanyFinanceDigest(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/fs/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
近一年各季度利润
*/
func GetCompanyProfit(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/pf/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
近一年各季度现金流
*/
func GetCompanyCashFlow(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/cf/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
近年业绩预告
*/
func GetCompanyPerformance(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/ep/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

/**
财务指标
*/
func GetCompanyFinanceIndex(code string) (response string, err error) {
	getCompanyShare := "http://ig507.com/data/time/f10/fi/" + code + "?licence=" + conf.Cnf.Licence
	return commonRequest(getCompanyShare)
}

func commonRequest(targetUrl string) (response string, err error) {
	Url, err := url.Parse(targetUrl)
	if err != nil {
		return
	}
	urlPath := Url.String()

	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	response = string(body)

	return response, err
}
