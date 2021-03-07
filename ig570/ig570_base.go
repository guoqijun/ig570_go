package ig570

import (
	"fmt"
	"ig570_go/conf"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetCompanyList() (response string, err error) {
	getCompanyUrl := "http://ig507.com/data/base/gplist"
	fmt.Print(getCompanyUrl)

	params := url.Values{}
	Url, err := url.Parse(getCompanyUrl)
	if err != nil {
		return
	}
	params.Set("licence", conf.Cnf.Licence)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	response = string(body)

	return response, err
}

func GetCompanyInfo(code string) (response string, err error) {
	getCompanyInfo := "http://ig507.com/data/time/f10/info/" + code + "?licence=" + conf.Cnf.Licence
	fmt.Print(getCompanyInfo)

	Url, err := url.Parse(getCompanyInfo)
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
