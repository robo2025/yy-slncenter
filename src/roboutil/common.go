package roboutil

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type userInfo struct {
	Msg     string            `json:"msg"`
	Data    map[string]string `json:"data"`
	Rescode string            `json:"rescode"`
}

func TimeToStamp(startTime,endTime string) (s,e int) {
	if startTime == "" && endTime == "" {
		e = int(time.Now().Unix())
		return s,e
	}
	ss, _ := time.Parse("2006-01-02 15:04:05", startTime)
	s = int(ss.Unix())

	ee, _ := time.Parse("2006-01-02 15:04:05", endTime)
	e = int(ee.Unix())

	return s,e
}
//   获取(用户名)
func HttpGet(UID int) string {
	url :=fmt.Sprintf("https://testapi.robo2025.com/user/service/usernames/%d" ,UID)
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if err != nil {
	}

	user := new(userInfo)
	err = json.Unmarshal(body, user)

	return user.Data["username"]
}
