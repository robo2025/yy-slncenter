package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"roboutil"
	"fmt"
	"strconv"
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
		return resp, nil
	}else {
		err := errors.New("获取详情时发生错误")
		return nil,err
	}
	errs := errors.New("获取详情时发生错误")
	return nil, errs
}

// 方案报价
func OfferSolution(db *gorm.DB, params *OfferParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	mainUserId := c.MustGet("main_user_id").(int)
	dbParams := prepareOfferData(params, uid, mainUserId)
	return writeOfferData(db, dbParams, uid)
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
	var userIds []int	// operator赋值
	for i:=0;i<len(offerOperation);i++ {
		userIds = append(userIds, offerOperation[i].OperatorId)
	}
	userMap := roboutil.HttpGetNames(userIds)

	for i:=0;i<len(offerOperation);i++ {
		offerOperation[i].Operator = fmt.Sprintf("供应商(%s)",userMap[strconv.Itoa(offerOperation[i].OperatorId)])
	}
	resp := offerOperation
	return resp, nil
}
