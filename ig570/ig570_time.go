package ig570

import (
	"ig570_go/common"
	"ig570_go/conf"
)

/**
描述：根据股票代码以及分时级别获取分时交易数据，交易时间从远到近排序。
目前 分时级别 支持5分钟、15分钟、30分钟、60分钟、日周月年级别（包括前后复权），
对应的值分别是 5、15、30、60、Day、Day_qfq（日线前复权）、Day_hfq（日线后复权）、
Week、Week_qfq（周线前复权）、Week_hfq（周线后复权）、Month、Month_qfq（月线前复权）、
Month_hfq（月线后复权）、Year、Year_qfq（年线前复权）、Year_hfq（年线后复权）
*/
func GetTimeRealLevel(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/time/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeRealKdj(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/kdj/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeRealMacd(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/macd/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeRealMa(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/ma/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}

func GetTimeRealBoll(code string, level string) (response string, err error) {
	requestUrl := "http://ig507.com/data/time/real/boll/" + code + "/" + level + "?licence=" + conf.Cnf.Licence
	return common.Request(requestUrl)
}
