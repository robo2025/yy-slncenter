package robodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InitDB(sqlURL string) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", sqlURL)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	return db, nil
}

func prepareSolutionData(params *CreateSolutionParams) *CreateSolutionParams {
	// 准备数据
	slnNo := params.SlnNo

	// basic info
	basicInfo := params.SlnBasicInfo
	basicInfo.SlnNo = slnNo
	basicInfo.SlnDate = time.Now()

	// user info
	userInfo := params.SlnUserInfo
	userInfo.SlnNo = slnNo

	// welding info
	weldingInfo := params.WeldingInfo
	weldingInfo.SlnNo = slnNo

	// device info
	deviceInfo := params.DeviceInfo
	if len(deviceInfo) != 0 {
		for _, el := range deviceInfo {
			el.SlnNo = slnNo
			el.SlnRole = "C"
		}
	}

	// device file
	deviceFile := make([]*WeldingFile, 0)

	// device image
	deviceImage := params.DeviceImg
	if len(deviceImage) != 0 {
		for _, el := range deviceImage {
			el.SlnNo = slnNo
			el.SlnRole = "C"
			el.FileType = "img"
			deviceFile = append(deviceFile, el)
		}
	}

	// device cad
	deviceCAD := params.DeviceCAD
	if deviceCAD != nil {
		deviceCAD.SlnNo = slnNo
		deviceCAD.SlnRole = "C"
		deviceCAD.FileType = "cad"
		deviceFile = append(deviceFile, deviceCAD)
	}

	// device attachment
	deviceAttachment := params.DeviceAttachment
	if deviceAttachment != nil {
		deviceAttachment.SlnNo = slnNo
		deviceAttachment.SlnRole = "C"
		deviceAttachment.FileType = "cad"
		deviceFile = append(deviceFile, deviceAttachment)
	}

	resp := &CreateSolutionParams{
		SlnBasicInfo: basicInfo,
		SlnUserInfo:  userInfo,
		WeldingInfo:  weldingInfo,
		DeviceInfo:   deviceInfo,
		DeviceFile:   deviceFile,
	}
	return resp
}

func writeSolutionData(db *gorm.DB, params *CreateSolutionParams) error {
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

	// 写入 sln_basic_info
	err = tx.Create(params.SlnBasicInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 sln_user_info
	err = tx.Create(params.SlnUserInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 welding_info
	err = tx.Create(params.WeldingInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 device_info
	if len(params.DeviceInfo) != 0 {
		for _, el := range params.DeviceInfo {
			err = tx.Create(el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 device_file
	if len(params.DeviceFile) != 0 {
		for _, el := range params.DeviceFile {
			err = tx.Create(el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
