package neo

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(apiUrl string, data url.Values, header map[string]string) (string, error) {
	httpClient := http.Client{}

	req, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	for key, value := range header {
		req.Header.Add(key, value)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		logs.Info(err.Error())
		return "", err
	}
	defer resp.Body.Close()

	result, bodyErr := ioutil.ReadAll(resp.Body)
	return string(result), bodyErr
}
