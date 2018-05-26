package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
)

// 获取方案列表
func FetchSolutionList(db *gorm.DB, uid int) ([]SlnBasicInfo, error) {
	dbData := []SlnBasicInfo{}
	db.Where("customer_id = ?", uid).Find(&dbData)
	if len(dbData) == 0 {
		return nil, errors.New("找不到方案列表")
	}
	return dbData, nil
}

// 获取方案细节
func FetchSolutionDetail(db *gorm.DB, slnID string, uid int) (*SolutionParams, error) {
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
func CreateSolution(db *gorm.DB, params *SolutionParams, uid int) error {
	dbParams := prepareSolutionData(params, uid)
	return writeSolutionData(db, dbParams)
}

// 更新现有方案
func UpdateSolution(db *gorm.DB, params *SolutionParams, uid int) error {
	dbParams := prepareSolutionData(params, uid)
	return updateSolutionData(db, dbParams)
}
