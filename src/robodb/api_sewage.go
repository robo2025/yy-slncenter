package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"fmt"
	"reflect"
)

//创建新污水方案
func CreateSewage(db *gorm.DB, params *SewageParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareSewageData(params, uid)
	return writeSewageData(db, dbParams)
}

// 获取污水方案细节
func FetchSewageDetail(db *gorm.DB, c *gin.Context) (*SewageDetailParams, error) {
	slnID := c.Param("id")

	var err error
	customer := &SewageParams{}
	supplier := &OfferParams{}
	resp := &SewageDetailParams{}

	// 读取用户询价数据
	customer, err = readSewageData(db, slnID, c)
	if err != nil {
		return nil, err
	}
	fmt.Println(reflect.TypeOf(resp))
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

func UpdateSewage(db *gorm.DB, params *SewageParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareSewageData(params, uid)
	return updateSewageData(db, dbParams)
}

