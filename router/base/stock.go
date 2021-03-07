package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ig570_go/common"
	"ig570_go/conf"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetCompany(ctx *gin.Context) {
	appG := common.Gin{C: ctx}

	getCompanyUrl := conf.ApiUrlCf.BaseGplist
	fmt.Print(getCompanyUrl)

	params := url.Values{}
	Url, err := url.Parse(getCompanyUrl)
	if err != nil {
		return
	}
	params.Set("licence", conf.ApiUrlCf.Licence)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	appG.Response(http.StatusOK, 0, string(body))
	return
}
