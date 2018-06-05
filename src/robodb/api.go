package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"strings"
	log "github.com/sirupsen/logrus"
)

// 获取方案列表
func FetchSolutionList(db *gorm.DB, c *gin.Context) ([]SlnBasicInfo, error) {
	uid := c.MustGet("uid").(int)
	is_type := c.Query("is_type")

	dbData := []SlnBasicInfo{}
	if is_type != "" && is_type != "all" {
		db.Where("customer_id = ? AND sln_status = ?", uid, strings.ToUpper(is_type)).Find(&dbData)
	} else {
		db.Where("customer_id = ?", uid).Find(&dbData)
	}

	if len(dbData) == 0 {
		log.Debug("找不到方案列表")
	}

	return dbData, nil
}

// 获取方案细节
func FetchSolutionDetail(db *gorm.DB, c *gin.Context) (*SolutionDetailParams, error) {
	slnID := c.Param("id")
	uid := c.MustGet("uid").(int)

	var err error
	customer := &SolutionParams{}
	supplier := &OfferParams{}
	resp := &SolutionDetailParams{}

	// 读取用户询价数据
	customer, err = readSolutionData(db, slnID, uid)
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
