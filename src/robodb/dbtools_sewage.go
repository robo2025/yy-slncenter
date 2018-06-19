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

	// SlnDevice 表
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
	resp := &SewageParams{
		SlnNo:        slnNo,
		UID:          uid,
		SlnBasicInfo: slnBasicInfo,
		SlnUserInfo:  slnUserInfo,
		SewageInfo:   sewageInfo,
		SlnDevice:    slnDevice,
		SlnFile:      slnFile,
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

	// 写入 sewage_info 表
	err = tx.Create(params.SewageInfo).Error
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

// 查询用户污水方案细节 new add sweage or welding
func readSewageData(db *gorm.DB, slnID string, c *gin.Context) (*SewageParams, error) {
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	SewageInfo := &SewageInfo{}
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
	db.Where("sln_no = ?", slnID).First(SewageInfo)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&slnDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, customerID).Find(&slnFile)

	resp := &SewageParams{
		SlnNo:        slnID,
		SlnBasicInfo: slnBasicInfo,
		SlnUserInfo:  slnUserInfo,
		SewageInfo:   SewageInfo,
		SlnDevice:    slnDevice,
		SlnFile:      slnFile,
	}
	return resp, nil
}

// 更新污水方案数据
func updateSewageData(db *gorm.DB, params *SewageParams) error {
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
