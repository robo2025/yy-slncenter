package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
)

func FetchSolutionList(db *gorm.DB) ([]SlnBasicInfo, error) {
	dbData := []SlnBasicInfo{}
	db.Find(&dbData)
	if len(dbData) == 0 {
		return nil, errors.New("找不到方案列表")
	}
	return dbData, nil
}

func FetchSolutionDetail(db *gorm.DB, slnID string) (*SlnBasicInfo, error) {
	dbData := &SlnBasicInfo{}
	db.Where("sln_no = ?", slnID).First(dbData)
	if dbData.ID == 0 {
		return nil, errors.New("找不到相应方案")
	}
	return dbData, nil
}

func CreateSolution(db *gorm.DB, params *SolutionParams) error {
	dbParams := prepareSolutionData(params)
	return writeSolutionData(db, dbParams)
}

func UpdateSolution(db *gorm.DB, params *SolutionParams) error {
	dbParams := prepareSolutionData(params)
	return updateSolutionData(db, dbParams)
}
