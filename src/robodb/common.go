package robodb

import (
	"time"
)

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
