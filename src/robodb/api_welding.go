package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
	"errors"
)

// 获取方案列表
func FetchSolutionList(db *gorm.DB, c *gin.Context) ([]SlnBasicInfo, error) {
	uid := c.MustGet("uid").(int)
	role := c.MustGet("role").(int)
	isType := c.Query("is_type")

	dbData := []SlnBasicInfo{}

	switch role {
	case 1: // customer
		if isType != "" && isType != "all" {
			db.Where("customer_id = ? AND sln_status = ?", uid, strings.ToUpper(isType)).Find(&dbData)
		} else {
			db.Where("customer_id = ?", uid).Find(&dbData)
		}

	case 2: // supplier
		// 只能查看已发布和已报价的
		db.Where("sln_status in (?)", []string{"P", "M"}).Find(&dbData)

	case 3, 4: // admin
		if isType != "" && isType != "all" {
			db.Where("sln_status = ?", strings.ToUpper(isType)).Find(&dbData)
		} else {
			db.Find(&dbData)
		}
	}

	return dbData, nil
}

// 获取方案细节
func FetchWeldingDetail(db *gorm.DB, c *gin.Context) (*WeldingDetailParams, error) {
	slnID := c.Param("id")

	var err error
	customer := &WeldingParams{}
	supplier := &OfferParams{}
	resp := &WeldingDetailParams{}

	// 读取用户询价数据
	customer, err = readSolutionData(db, slnID, c)
	if err != nil {
		return nil, err
	}
	resp.Customer = customer

	// 读取报价数据
	resp.Supplier = nil
	if customer.SlnBasicInfo.SlnStatus == string(SlnStatusOffer) {
		supplier, err = readOfferData(db, slnID, customer.SlnBasicInfo.SupplierID)
		if err == nil {
			resp.Supplier = supplier
		}
	}

	return resp, nil
}

// 创建新焊接方案
func CreateSolution(db *gorm.DB, params *WeldingParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareWeldingData(params, uid)
	return writeWeldingData(db, dbParams)
}

// 更新现有焊接方案
func UpdateWelding(db *gorm.DB, params *WeldingParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareWeldingData(params, uid)
	return updateWeldingData(db, dbParams)
}

// 方案报价
func OfferSolution(db *gorm.DB, params *OfferParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareOfferData(params, uid)
	return writeOfferData(db, dbParams)
}

// RPC 查询
func FetchSolutionRPC(db *gorm.DB, params *SolutionRPCReqParams) (map[string]interface{}, error) {

	uid := params.UID
	resp := make(map[string]interface{})
	respFlag := true

	for _, el := range params.Solution {
		solutionRPC, err := readSolutionRPCData(db, el, uid)
		if err != nil {
			resp[el] = &SolutionRPCParams{
				Success:  false,
				ErrorMsg: err.Error(),
			}
			respFlag = false
		} else {
			resp[el] = solutionRPC
		}
	}

	if respFlag {
		return resp, nil
	}

	return resp, errors.New("部分方案状态错误")
}

// RPC 查询方案细节
func FetchSolutionRPCDetail(db *gorm.DB, c *gin.Context) (*SolutionRPCParams, error) {
	slnID := c.Param("id")
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		return nil, err
	}
	return readSolutionRPCData(db, slnID, uid)
}
