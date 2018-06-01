package roboweb

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func HttpRequest(method string, url string, auth string) ([]byte, error) {

	// 准备请求 Client
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", auth)

	// 请求数据
	response, _ := client.Do(request)
	defer response.Body.Close()

	// 解析数据
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return body, nil
	} else {
		return nil, fmt.Errorf("Response error: %d", response.StatusCode)
	}
}
