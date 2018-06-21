package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

func FetchDetail(db *gorm.DB, c *gin.Context) (interface{}, error) {
	slnID := c.Param("id")
	basicInfo := &SlnBasicInfo{}
	db.Where("sln_no = ?", slnID).First(basicInfo)
	if basicInfo.SlnType == "welding" {
		resp, _ := FetchWeldingDetail(db, c)

		return resp, nil
	} else if basicInfo.SlnType == "sewage" {
		resp, _ := FetchSewageDetail(db, c)

		return resp, nil
	}
	return nil, nil
}