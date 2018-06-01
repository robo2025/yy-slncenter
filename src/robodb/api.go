package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
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
		return nil, errors.New("找不到方案列表")
	}
	return dbData, nil
}

// 获取方案细节
func FetchSolutionDetail(db *gorm.DB, c *gin.Context) (*SolutionParams, error) {
	slnID := c.Param("id")
	uid := c.MustGet("uid").(int)

	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	WeldingInfo := &WeldingInfo{}
	weldingDevice := []WeldingDevice{}
	weldingFile := []WeldingFile{}

	db.Where("sln_no = ? AND customer_id = ?", slnID, uid).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return nil, errors.New("找不到相应方案")
	}
	db.Where("sln_no = ?", slnID).First(slnUserInfo)
	db.Where("sln_no = ?", slnID).First(WeldingInfo)
	db.Where("sln_no = ? AND sln_role = ?", slnID, "C").Find(&weldingDevice)
	db.Where("sln_no = ? AND sln_role = ?", slnID, "C").Find(&weldingFile)

	resp := &SolutionParams{
		SlnNo:         slnID,
		SlnBasicInfo:  slnBasicInfo,
		SlnUserInfo:   slnUserInfo,
		WeldingInfo:   WeldingInfo,
		WeldingDevice: weldingDevice,
		WeldingFile:   weldingFile,
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
