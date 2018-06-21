package robodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"errors"
	"strings"
	"github.com/gin-gonic/gin"
)

// InitDB 初始化数据库
func InitDB(sqlURL string) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	return db, nil
}

// 准备写入或者更新用的方案数据
func prepareWeldingData(params *WeldingParams, uid int) *WeldingParams {
	// 准备数据
	slnNo := strings.TrimSpace(params.SlnNo)

	params.UID = uid

	// sln_basic_info 表
	slnBasicInfo := params.SlnBasicInfo
	currentDate := time.Now()
	if slnBasicInfo != nil {
		slnBasicInfo.SlnNo = slnNo
		slnBasicInfo.CustomerID = uid
		slnBasicInfo.SlnDate = int(currentDate.Unix())
		slnBasicInfo.SlnExpired = int(currentDate.AddDate(0, 0, 90).Unix())
	}

	// sln_user_info 表
	slnUserInfo := params.SlnUserInfo
	if slnUserInfo != nil {
		slnUserInfo.SlnNo = slnNo
	}

	// welding_info 表
	weldingInfo := params.WeldingInfo
	if weldingInfo != nil {
		weldingInfo.SlnNo = slnNo
	}

	// sln_device 表
	slnDevice := make([]SlnDevice, 0)
	if len(params.SlnDevice) != 0 {
		for _, el := range params.SlnDevice {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			slnDevice = append(slnDevice, el)
		}
	}

	// sln_file 表
	slnFile := make([]SlnFile, 0)
	if len(params.SlnFile) != 0 {
		for _, el := range params.SlnFile {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			slnFile = append(slnFile, el)
		}
	}

	// 返回组合数据
	resp := &WeldingParams{
		SlnNo:        slnNo,
		UID:          uid,
		SlnBasicInfo: slnBasicInfo,
		SlnUserInfo:  slnUserInfo,
		WeldingInfo:  weldingInfo,
		SlnDevice:    slnDevice,
		SlnFile:      slnFile,
	}
	return resp
}

// 准备方案报价数据
func prepareOfferData(params *OfferParams, uid int) *OfferParams {
	// 准备数据
	slnNo := strings.TrimSpace(params.SlnNo)

	params.UID = uid

	// sln_supplier_info
	slnSupplierInfo := params.SlnSupplierInfo
	currentDate := time.Now()
	if slnSupplierInfo != nil {
		slnSupplierInfo.SlnNo = slnNo
		slnSupplierInfo.UserID = uid
		slnSupplierInfo.ExpiredDate = int(currentDate.AddDate(0, 0, 30).Unix())
	}

	// welding_device 表
	slnDevice := make([]SlnDevice, 0)
	if len(params.SlnDevice) != 0 {
		for _, el := range params.SlnDevice {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "S"
			slnDevice = append(slnDevice, el)
		}
	}

	// sln_support 表
	slnSupport := make([]SlnSupport, 0)
	if len(params.SlnSupport) != 0 {
		for _, el := range params.SlnSupport {
			el.SlnNo = slnNo
			el.UserID = uid
			slnSupport = append(slnSupport, el)
		}
	}

	// welding_tech_param 表
	weldingTechParam := make([]WeldingTechParam, 0)
	if len(params.WeldingTechParam) != 0 {
		for _, el := range params.WeldingTechParam {
			el.SlnNo = slnNo
			el.UserID = uid
			weldingTechParam = append(weldingTechParam, el)
		}
	}

	// sln_file 表
	slnFile := make([]SlnFile, 0)
	if len(params.SlnFile) != 0 {
		for _, el := range params.SlnFile {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			slnFile = append(slnFile, el)
		}
	}

	// 返回组合数据
	resp := &OfferParams{
		SlnNo:            slnNo,
		UID:              uid,
		SlnSupplierInfo:  slnSupplierInfo,
		SlnDevice:        slnDevice,
		SlnSupport:       slnSupport,
		WeldingTechParam: weldingTechParam,
		SlnFile:          slnFile,
	}
	return resp
}

// 写入方案数据 writeSolutionData
func writeWeldingData(db *gorm.DB, params *WeldingParams) error {
	var err error

	// 写入数据库
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 写入 sln_basic_info 表
	err = tx.Create(params.SlnBasicInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 sln_user_info 表
	err = tx.Create(params.SlnUserInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 welding_info 表
	err = tx.Create(params.WeldingInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 Sln_device 表
	if len(params.SlnDevice) != 0 {
		for _, el := range params.SlnDevice {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 sln_file 表
	if len(params.SlnFile) != 0 {
		for _, el := range params.SlnFile {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 更新方案数据 updateSolutionData
func updateWeldingData(db *gorm.DB, params *WeldingParams) error {
	var err error
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	slnFile := []SlnFile{}

	// 查找数据库相应数据
	slnNo := params.SlnNo
	uid := params.UID
	db.Where("sln_no = ? AND customer_id = ?", slnNo, uid).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return errors.New("找不到相应方案")
	}
	db.Where("sln_no = ?", slnNo).First(slnUserInfo)
	db.Where("sln_no = ? AND sln_role = ?", slnNo, "C").Find(&slnFile)

	// 写入数据库
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 更新 sln_basic_info 表
	if slnBasicInfo != nil && params.SlnBasicInfo != nil {
		params.SlnBasicInfo.ID = slnBasicInfo.ID
		err = tx.Save(params.SlnBasicInfo).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新 sln_user_info 表
	if slnUserInfo != nil && params.SlnUserInfo != nil {
		params.SlnUserInfo.ID = slnUserInfo.ID
		err = tx.Save(params.SlnUserInfo).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新 sln_file 表
	if len(params.SlnFile) != 0 {

		// 删除所有的旧数据
		err = db.Where("sln_no = ? AND sln_role = ?", slnNo, "C").Delete(SlnFile{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// 插入所有的新数据
		for _, el := range params.SlnFile {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 写入方案数据
func writeOfferData(db *gorm.DB, params *OfferParams) error {
	var err error

	slnBasicInfo := &SlnBasicInfo{}
	db.Where("sln_no = ?", params.SlnNo).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return errors.New("找不到相应方案")
	}

	if slnBasicInfo.SlnStatus != string(SlnStatusPublish) {
		return errors.New("该方案不可以报价")
	}

	// 写入数据库
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 更新 sln_basic_info 表
	tx.Model(slnBasicInfo).Updates(SlnBasicInfo{
		SupplierID:    params.SlnSupplierInfo.UserID,
		SupplierPrice: params.SlnSupplierInfo.TotalPrice,
		SlnStatus:     string(SlnStatusOffer),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 sln_supplier_info 表
	err = tx.Create(params.SlnSupplierInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 sln_device 表
	if len(params.SlnDevice) != 0 {
		for _, el := range params.SlnDevice {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_support 表
	if len(params.SlnSupport) != 0 {
		for _, el := range params.SlnSupport {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_tech_param 表
	if len(params.WeldingTechParam) != 0 {
		for _, el := range params.WeldingTechParam {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_file 表
	if len(params.SlnFile) != 0 {
		for _, el := range params.SlnFile {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 查询用户方案细节
func readSolutionData(db *gorm.DB, slnID string, c *gin.Context) (*WeldingParams, error) {
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	WeldingInfo := &WeldingInfo{}
	slnDevice := []SlnDevice{}
	slnFile := []SlnFile{}

	uid := c.MustGet("uid").(int)
	role := c.MustGet("role").(int)
	switch role {
	case 1: // customer
		db.Where("sln_no = ? AND customer_id = ?", slnID, uid).First(slnBasicInfo)
	case 2, 3, 4: // supplier, admin, super
		db.Where("sln_no = ?", slnID).First(slnBasicInfo)
	}

	if slnBasicInfo.ID == 0 {
		return nil, errors.New("找不到相应方案")
	}

	customerID := slnBasicInfo.CustomerID
	db.Where("sln_no = ?", slnID).First(slnUserInfo)
	db.Where("sln_no = ?", slnID).First(WeldingInfo)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&slnDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&slnFile)

	resp := &WeldingParams{
		SlnNo:        slnID,
		SlnBasicInfo: slnBasicInfo,
		SlnUserInfo:  slnUserInfo,
		WeldingInfo:  WeldingInfo,
		SlnDevice:    slnDevice,
		SlnFile:      slnFile,
	}

	return resp, nil
}

// 查询供应商报价细节
func readOfferData(db *gorm.DB, slnID string, uid int) (*OfferParams, error) {
	slnSupplierInfo := &SlnSupplierInfo{}
	slnDevice := []SlnDevice{}
	weldingTechParams := []WeldingTechParam{}
	slnSupport := []SlnSupport{}
	slnFile := []SlnFile{}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).First(slnSupplierInfo)
	if slnSupplierInfo.ID == 0 {
		return nil, errors.New("找不到相应报价")
	}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingTechParams)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnSupport)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnFile)

	resp := &OfferParams{
		SlnNo:            slnID,
		SlnSupplierInfo:  slnSupplierInfo,
		SlnDevice:        slnDevice,
		WeldingTechParam: weldingTechParams,
		SlnSupport:       slnSupport,
		SlnFile:          slnFile,
	}

	return resp, nil
}

// 提供 RPC 查询
func readSolutionRPCData(db *gorm.DB, slnID string, uid int) (*SolutionRPCParams, error) {
	slnBasicInfo := &SlnBasicInfo{}
	db.Where("sln_no = ? AND customer_id = ?", slnID, uid).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return nil, errors.New("找不到相应方案")
	}

	if slnBasicInfo.SlnStatus != string(SlnStatusOffer) {
		return nil, errors.New("该方案不属于已报价状态")
	}

	slnSupplierInfo := &SlnSupplierInfo{}
	db.Where("sln_no = ? AND user_id = ?", slnID, slnBasicInfo.SupplierID).First(slnSupplierInfo)
	if slnSupplierInfo.ID == 0 {
		return nil, errors.New("找不到相应报价")
	}

	resp := &SolutionRPCParams{
		SlnBasicInfo:    slnBasicInfo,
		SlnSupplierInfo: slnSupplierInfo,
		Success:         true,
	}

	return resp, nil
}
