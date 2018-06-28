package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"fmt"
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
		fmt.Println(resp)
		fmt.Println(resp.Customer.SewageInfo)
		return resp, nil
	}
	return nil, nil
}

// 方案指派
func AssignSolution(db *gorm.DB, params *AssignParams, c *gin.Context) error {
	dbParams := prepareAssignData(params)
	return writeAssignData(db, dbParams)
}