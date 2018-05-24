package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
)

func FetchSolutionList(db *gorm.DB) ([]SlnBasicInfo, error) {
	dbData := []SlnBasicInfo{}
	db.Find(&dbData)
	if len(dbData) == 0 {
		return nil, errors.New("no matched data")
	}
	return dbData, nil
}

func FetchSolutionDetail(db *gorm.DB, slnID string) (SlnBasicInfo, error) {
	dbData := SlnBasicInfo{}
	db.Where("sln_no = ?", slnID).First(&dbData)
	if dbData.ID == 0 {
		return dbData, errors.New("no matched data")
	}
	return dbData, nil
}

func CreateSolution(db *gorm.DB, params *CreateSolutionParams) error {
	dbParams := prepareSolutionData(params)
	return writeSolutionData(db, dbParams)
}
