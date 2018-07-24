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

type UsersInfo struct {
	Rescode string            `json:"rescode"`
	Data    map[string]string `json:"data"`
	Msg     string            `json:"msg"`
}

func TimeToStamp(startTime,endTime string) (s,e int) {
	if startTime == "" && endTime == "" {
		e = int(time.Now().Unix())
		return s,e
	}
	startTime += " 00:00:00"
	endTime += " 23:59:59"

	ss, _ := time.Parse("2006-01-02 15:04:05", startTime)
	s = int(ss.Unix())

	ee, _ := time.Parse("2006-01-02 15:04:05", endTime)
	e = int(ee.Unix())

	if s > e{
		return e,s
	}
	return s,e
}
func TimeToStamp2(startTime,endTime string) (s,e int) {
	if startTime == "" && endTime == "" {
		return s,e
	}
	startTime += " 00:00:00"
	endTime += " 23:59:59"

	ss, _ := time.Parse("2006-01-02 15:04:05", startTime)
	s = int(ss.Unix())

	ee, _ := time.Parse("2006-01-02 15:04:05", endTime)
	e = int(ee.Unix())

	if s > e{
		return e,s
	}
	return s,e
}
//   获取(用户名)单个
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

//   获取(用户名)多个
func HttpGetNames(UIDs []int) map[string]string {
	url := "https://testapi.robo2025.com/user/service/usernames?"
	for i := 0; i < len(UIDs); i++ {
		url += fmt.Sprintf("user_ids=%d&", UIDs[i])
	}
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if err != nil {
	}
	users := new(UsersInfo)
	err = json.Unmarshal(body,users)

	return users.Data
}

// 列表去重
func LikeSetFromPy(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{}
	for _, e := range slc{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{
			result = append(result, e)
		}
	}
	return result
}