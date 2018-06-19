package robodb

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"fmt"
	"reflect"
)

type empty1 interface {
}

//创建新污水方案
func CreateSewage(db *gorm.DB, params *SewageParams, c *gin.Context) error {
	uid := c.MustGet("uid").(int)
	dbParams := prepareSewageData(params, uid)
	return writeSewageData(db, dbParams)
}

// 获取污水方案细节
//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
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

//func jugeSlnType(db *gorm.DB, id int) (empty1,empty1){
//	slnBasicInfo := &SlnBasicInfo{}
//	db.Where("sln_no = ?", id).First(slnBasicInfo)
//
//	var (
//		customer interface{}
//		resp interface{}
//	)
//	if slnBasicInfo.SlnType == "焊接" {
//		customer = &SolutionParams{}
//		resp = &SolutionDetailParams{}
//
//	} else if slnBasicInfo.SlnType == "污水处理" {
//		customer = &SewageParams{}
//		resp = &SewageDetailParams{}
//	}
//	return customer,resp
//}
