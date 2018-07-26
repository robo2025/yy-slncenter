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

type AllUsersInfo struct {
	Rescode string                   `json:"rescode"`
	Data    []map[string]interface{} `json:"data"`
	Msg     string                   `json:"msg"`
}

func StartTimeToStamp(startTime string) int {
	if startTime == "" {
		return 0
	}
	startTime += " 00:00:00"
	s, _ := time.Parse("2006-01-02 15:04:05", startTime)
	return int(s.Unix())
}

func EndTimeToStamp(endTime string) int {
	if endTime == "" {
		return 0
	}
	endTime += " 23:59:59"
	s, _ := time.Parse("2006-01-02 15:04:05", endTime)
	return int(s.Unix())
}

func TimeRangeToStamp(startTime, endTime string) (s, e int) {
	if startTime == "" && endTime == "" {
		e = int(time.Now().Unix())
		return s, e
	}
	startTime += " 00:00:00"
	endTime += " 23:59:59"

	ss, _ := time.Parse("2006-01-02 15:04:05", startTime)
	s = int(ss.Unix())

	ee, _ := time.Parse("2006-01-02 15:04:05", endTime)
	e = int(ee.Unix())

	if s > e {
		return e, s
	}
	return s, e
}

//   获取(用户名)单个
func HttpGet(UID int) string {
	url := fmt.Sprintf("https://testapi.robo2025.com/user/service/usernames/%d", UID)
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
	err = json.Unmarshal(body, users)

	return users.Data
}

// 列表去重
func LikeSetFromPy(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// 根据名字获取ID https://testapi.robo2025.com/user/service/users/all
func HttpGetId(name string) []int {
	var result []int
	url := "https://testapi.robo2025.com/user/service/users/all?"
	url += fmt.Sprintf("username=%s&", name)

	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //请求数据进行读取
	if err != nil {
	}
	users := new(AllUsersInfo)
	err = json.Unmarshal(body, users)
	for _, v := range users.Data {
		resultId, _ := v["main_user_id"].(float64)
		result = append(result, int(resultId))
	}

	return result
}
