package roboweb

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	"time"
)

// Refer: https://medium.com/myntra-engineering/my-journey-with-golang-web-services-4d922a8c9897
func HttpRequest(method string, url string, header map[string]string, body io.Reader) ([]byte, error) {
	// 创建自定义 Client
	// 设置超时为 10 秒
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Duration(10 * time.Second),
	}

	// 创建 NewRequest
	req, _ := http.NewRequest(method, url, body)

	// Add headers
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// 进行网络请求
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 解析返回数据
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response error: %d", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}
