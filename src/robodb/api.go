package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"fmt"
)
// 可获取多种方案详情
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

// 方案报价
func OfferSolution(db *gorm.DB, params *OfferParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareOfferData(params, uid)
	return writeOfferData(db, dbParams)
}

// 方案指派
func AssignSolution(db *gorm.DB, params *AssignParams, c *gin.Context) error {
	dbParams := prepareAssignData(params)
	return writeAssignData(db, dbParams)
}

// 获取操作记录
func FetchLog(db *gorm.DB, c *gin.Context) ([]OperationLog, error) {
	slnNo := c.Query("sln_no")

	operationLog := []OperationLog{}
	db.Order("-add_time").Where("sln_no = ?", slnNo).Find(&operationLog)

	resp := operationLog
	return resp, nil
}

//获取报价操作记录
func FetchOfferOperation(db *gorm.DB, c *gin.Context) ([]OfferOperation, error) {
	slnNo := c.Query("sln_no")
	sbmNo := c.Query("sbm_no")

	offerOperation := []OfferOperation{}
	db.Order("add_time").Where("sln_no = ? And sbm_no = ?", slnNo, sbmNo).Find(&offerOperation)

	resp := offerOperation
	return resp, nil
}
