package ig570

import (
	"ig570_go/common"
	"ig570_go/conf"
)

/**
实时交易数据接口
*/
func GetTimeReal(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
买卖五档盘口
*/
func GetTimeRealLevel5(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/trace/level5/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
当天逐笔交易
*/
func GetTimeRealOneByOne(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/trace/onebyone/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
当天分时成交
*/
func GetTimeRealTimeDeal(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/trace/timedeal/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

/**
当天逐笔大单交易
*/
func GetTimeRealBigDeal(code string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/trace/bigdeal/" + code + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}
