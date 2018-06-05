package robodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"errors"
	"strings"
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
func prepareSolutionData(params *SolutionParams, uid int) *SolutionParams {
	// 准备数据
	slnNo := strings.TrimSpace(params.SlnNo)

	params.UID = uid

	// sln_basic_info 表
	slnBasicInfo := params.SlnBasicInfo
	currentDate := time.Now()
	if slnBasicInfo != nil {
		slnBasicInfo.SlnNo = slnNo
		slnBasicInfo.CustomerID = uid
		slnBasicInfo.SlnDate = currentDate
		slnBasicInfo.SlnExpired = currentDate.AddDate(0, 0, 90)
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

	// welding_device 表
	weldingDevice := make([]WeldingDevice, 0)
	if len(params.WeldingDevice) != 0 {
		for _, el := range params.WeldingDevice {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			weldingDevice = append(weldingDevice, el)
		}
	}

	// welding_file 表
	weldingFile := make([]WeldingFile, 0)
	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			weldingFile = append(weldingFile, el)
		}
	}

	// 返回组合数据
	resp := &SolutionParams{
		SlnNo:         slnNo,
		UID:           uid,
		SlnBasicInfo:  slnBasicInfo,
		SlnUserInfo:   slnUserInfo,
		WeldingInfo:   weldingInfo,
		WeldingDevice: weldingDevice,
		WeldingFile:   weldingFile,
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
		slnSupplierInfo.ExpiredDate = currentDate.AddDate(0, 0, 30)
	}

	// welding_device 表
	weldingDevice := make([]WeldingDevice, 0)
	if len(params.WeldingDevice) != 0 {
		for _, el := range params.WeldingDevice {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "S"
			weldingDevice = append(weldingDevice, el)
		}
	}

	// welding_support 表
	weldingSupport := make([]WeldingSupport, 0)
	if len(params.WeldingSupport) != 0 {
		for _, el := range params.WeldingSupport {
			el.SlnNo = slnNo
			el.UserID = uid
			weldingSupport = append(weldingSupport, el)
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

	// welding_file 表
	weldingFile := make([]WeldingFile, 0)
	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "C"
			weldingFile = append(weldingFile, el)
		}
	}

	// 返回组合数据
	resp := &OfferParams{
		SlnNo:            slnNo,
		UID:              uid,
		SlnSupplierInfo:  slnSupplierInfo,
		WeldingDevice:    weldingDevice,
		WeldingSupport:   weldingSupport,
		WeldingTechParam: weldingTechParam,
		WeldingFile:      weldingFile,
	}
	return resp
}

// 写入方案数据
func writeSolutionData(db *gorm.DB, params *SolutionParams) error {
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

	// 写入 welding_device 表
	if len(params.WeldingDevice) != 0 {
		for _, el := range params.WeldingDevice {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_file 表
	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 更新方案数据
func updateSolutionData(db *gorm.DB, params *SolutionParams) error {
	var err error
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	weldingFile := []WeldingFile{}

	// 查找数据库相应数据
	slnNo := params.SlnNo
	uid := params.UID
	db.Where("sln_no = ? AND customer_id = ?", slnNo, uid).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return errors.New("找不到相应方案")
	}
	db.Where("sln_no = ?", slnNo).First(slnUserInfo)
	db.Where("sln_no = ? AND sln_role = ?", slnNo, "C").Find(&weldingFile)

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

	// 更新 welding_file 表
	if len(params.WeldingFile) != 0 {

		// 删除所有的旧数据
		err = db.Where("sln_no = ? AND sln_role = ?", slnNo, "C").Delete(WeldingFile{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// 插入所有的新数据
		for _, el := range params.WeldingFile {
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

	// 写入 welding_device 表
	if len(params.WeldingDevice) != 0 {
		for _, el := range params.WeldingDevice {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_support 表
	if len(params.WeldingSupport) != 0 {
		for _, el := range params.WeldingSupport {
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
	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
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
func readSolutionData(db *gorm.DB, slnID string, uid int) (*SolutionParams, error) {
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
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingFile)

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

// 查询供应商报价细节
func readOfferData(db *gorm.DB, slnID string, uid int) (*OfferParams, error) {
	slnSupplierInfo := &SlnSupplierInfo{}
	weldingDevice := []WeldingDevice{}
	weldingTechParams := []WeldingTechParam{}
	weldingSupport := []WeldingSupport{}
	weldingFile := []WeldingFile{}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).First(slnSupplierInfo)
	if slnSupplierInfo.ID == 0 {
		return nil, errors.New("找不到相应报价")
	}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingTechParams)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingSupport)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingFile)

	resp := &OfferParams{
		SlnNo:            slnID,
		SlnSupplierInfo:  slnSupplierInfo,
		WeldingDevice:    weldingDevice,
		WeldingTechParam: weldingTechParams,
		WeldingSupport:   weldingSupport,
		WeldingFile:      weldingFile,
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
	}

	return resp, nil
}
