package request

import (
	"net/http"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"strings"
	"fmt"
	"encoding/json"
	"time"
	"github.com/shawflying/beego-common-utils/utils/comutil"
)

func getResponseTime(start time.Time) time.Duration {
	end := time.Now()
	return end.Sub(start)
}

// Get 请求
func Get(url string) (content []byte, err error) {
	start := time.Now()
	logs.Info("request-get-url: " + url)
	req, err := http.Get(url)
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	logs.Info("request-get-data: " + comutil.TransInterfaceToString(body))
	logs.Info("request-get-response-time:", getResponseTime(start))
	return body, err
}

func Post(url string, params interface{}) (content []byte, err error) {
	start := time.Now()
	b, err := json.Marshal(params)
	logs.Info("request-post-url: " + url)
	logs.Info("request-post-params: " + comutil.TransInterfaceToString(b))
	req, err := http.Post(url, "application/json", strings.NewReader(string(b)))

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	logs.Info("request-post-data: " + comutil.TransInterfaceToString(body))
	logs.Info("request-post-response-time:", getResponseTime(start))
	return body, err
}

func PostForm(url string, params interface{}) (content []byte, err error) {
	start := time.Now()
	input, err := json.Marshal(params)

	mapParams := make(map[string]interface{});
	json.Unmarshal(input, &mapParams)

	paramsData := ""
	for k, v := range mapParams {
		paramsData += "&" + k + "=" + comutil.TransInterfaceToString(v)
	}
	fmt.Println(paramsData)
	logs.Info("request-postForm-url: " + url)
	logs.Info("request-postForm-params: ", params)
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(string(paramsData)))

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	logs.Info("request-postForm-data: " + comutil.TransInterfaceToString(body))
	logs.Info("request-postForm-response-time:", getResponseTime(start))
	return body, err
}
