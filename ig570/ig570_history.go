package ig570

import (
	"ig570_go/common"
	"ig570_go/conf"
)

func GetTimeHistoryTrade(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/history/trade/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeHistoryKdj(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/history/kdj/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeHistoryMacd(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/history/macd/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeHistoryMa(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/history/ma/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}
func GetTimeHistoryBoll(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/history/boll/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}
