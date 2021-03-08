package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Request(targetUrl string) (response string, err error) {
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
