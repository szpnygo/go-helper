package neo

import (
	"bytes"
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

func PostRaw(apiUrl string, data []byte, rawType string, header map[string]string) (string, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/"+rawType)
	for key, value := range header {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	msgResult, rErr := ioutil.ReadAll(resp.Body)
	return string(msgResult), rErr
}

func Get(apiUrl string, header map[string]string) (string, error) {
	httpClient := http.Client{}

	req, _ := http.NewRequest("GET", apiUrl, nil)
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
