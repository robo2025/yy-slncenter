package robodb

import (
	"strings"
	"time"
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/gin-gonic/gin"
)

// 准备写入或者更新用的方案数据
func prepareSewageData(params *SewageParams, uid int) *SewageParams {
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

	// sewage_info 表
	sewageInfo := params.SewageInfo
	if sewageInfo != nil {
		sewageInfo.SlnNo = slnNo
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
	resp := &SewageParams{
		SlnNo:         slnNo,
		UID:           uid,
		SlnBasicInfo:  slnBasicInfo,
		SlnUserInfo:   slnUserInfo,
		SewageInfo:    sewageInfo,
		WeldingDevice: weldingDevice,
		WeldingFile:   weldingFile,
	}
	return resp
}

// 写入污水方案数据
func writeSewageData(db *gorm.DB, params *SewageParams) error {
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
	err = tx.Create(params.SewageInfo).Error
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

// 查询用户污水方案细节 new add sweage or welding
func readSewageData(db *gorm.DB, slnID string, c *gin.Context) (*SewageParams, error) {
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	SewageInfo := &SewageInfo{}
	//WeldingInfo := &WeldingInfo{}
	weldingDevice := []WeldingDevice{}
	weldingFile := []WeldingFile{}

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
	//switch slnBasicInfo.SlnType {
	//case "焊接":
	//	db.Where("sln_no = ?", slnID).First(WeldingInfo)
	//case "污水处理":
	//	db.Where("sln_no = ?", slnID).First(SewageInfo)
	//}

	db.Where("sln_no = ?", slnID).First(slnUserInfo)
	db.Where("sln_no = ?", slnID).First(SewageInfo)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&weldingDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&weldingFile)

	//switch slnBasicInfo.SlnType {
	//case "焊接":
	//	resp := &SolutionParams{
	//		SlnNo:         slnID,
	//		SlnBasicInfo:  slnBasicInfo,
	//		SlnUserInfo:   slnUserInfo,
	//		WeldingInfo:   WeldingInfo,
	//		WeldingDevice: weldingDevice,
	//		WeldingFile:   weldingFile,
	//	}
	//	return resp, nil
	//case "污水处理":
	//	resp := &SewageParams{
	//		SlnNo:         slnID,
	//		SlnBasicInfo:  slnBasicInfo,
	//		SlnUserInfo:   slnUserInfo,
	//		SewageInfo:    SewageInfo,
	//		WeldingDevice: weldingDevice,
	//		WeldingFile:   weldingFile,
	//	}
	//	return resp, nil
	//}
	//if slnBasicInfo.SlnType == "焊接" {
	//	resp := &SolutionParams{
	//		SlnNo:         slnID,
	//		SlnBasicInfo:  slnBasicInfo,
	//		SlnUserInfo:   slnUserInfo,
	//		WeldingInfo:   WeldingInfo,
	//		WeldingDevice: weldingDevice,
	//		WeldingFile:   weldingFile,
	//	}
	//	return resp, nil
	//} else if slnBasicInfo.SlnType == "污水处理" {
	//	resp := &SewageParams{
	//		SlnNo:         slnID,
	//		SlnBasicInfo:  slnBasicInfo,
	//		SlnUserInfo:   slnUserInfo,
	//		SewageInfo:    SewageInfo,
	//		WeldingDevice: weldingDevice,
	//		WeldingFile:   weldingFile,
	//	}
	//	return resp, nil
	//}
	resp := &SewageParams{
		SlnNo:         slnID,
		SlnBasicInfo:  slnBasicInfo,
		SlnUserInfo:   slnUserInfo,
		SewageInfo:   SewageInfo,
		WeldingDevice: weldingDevice,
		WeldingFile:   weldingFile,
	}
	return resp,nil
}

// 更新污水方案数据   和焊接一模一样 后期优化必看
func updateSewageData(db *gorm.DB, params *SewageParams) error {
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