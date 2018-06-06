package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
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

	case 2, 3, 4: // supplier
		if isType != "" && isType != "all" {
			db.Where("sln_status = ?", strings.ToUpper(isType)).Find(&dbData)
		} else {
			db.Find(&dbData)
		}
	}

	return dbData, nil
}

// 获取方案细节
func FetchSolutionDetail(db *gorm.DB, c *gin.Context) (*SolutionDetailParams, error) {
	slnID := c.Param("id")

	var err error
	customer := &SolutionParams{}
	supplier := &OfferParams{}
	resp := &SolutionDetailParams{}

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

// 创建新方案
func CreateSolution(db *gorm.DB, params *SolutionParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareSolutionData(params, uid)
	return writeSolutionData(db, dbParams)
}

// 更新现有方案
func UpdateSolution(db *gorm.DB, params *SolutionParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareSolutionData(params, uid)
	return updateSolutionData(db, dbParams)
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

	for _, el := range params.Solution {
		solutionRPC, err := readSolutionRPCData(db, el, uid)
		if err != nil {
			resp[el] = &SolutionRPCParams{
				Success:  false,
				ErrorMsg: err.Error(),
			}
		} else {
			resp[el] = solutionRPC
		}
	}

	return resp, nil
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
