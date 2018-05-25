package robodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"errors"
	"strings"
)

func InitDB(sqlURL string) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	return db, nil
}

func prepareSolutionData(params *SolutionParams, uid int) *SolutionParams {
	// 准备数据
	slnNo := strings.TrimSpace(params.SlnNo)

	// sln_basic_info 表
	slnBasicInfo := params.SlnBasicInfo
	if slnBasicInfo != nil {
		slnBasicInfo.SlnNo = slnNo
		slnBasicInfo.CustomerID = uid
		slnBasicInfo.SlnDate = time.Now()
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
	weldingDevice := params.WeldingDevice
	if len(weldingDevice) != 0 {
		for _, el := range weldingDevice {
			el.SlnNo = slnNo
			el.SlnRole = "C"
		}
	}

	// welding_file 表
	weldingFile := make([]*WeldingFile, 0)

	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
			el.SlnNo = slnNo
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
			err = tx.Create(el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_file 表
	if len(params.WeldingFile) != 0 {
		for _, el := range params.WeldingFile {
			err = tx.Create(el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func updateSolutionData(db *gorm.DB, params *SolutionParams) error {
	var err error
	slnBasicInfo := &SlnBasicInfo{}
	slnUserInfo := &SlnUserInfo{}
	weldingFile := []WeldingFile{}

	// 查找数据库相应数据
	slnNo := params.SlnNo
	uid := params.UID
	db.Where("sln_no = ? AND customer_id = ?", slnNo, uid).First(slnBasicInfo)
	if slnBasicInfo == nil {
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
