package robodb

import (
	"database/sql"
)

func FetchSolutionList(db *sql.DB) ([]interface{}, error) {
	dbData, err := fetchMultiData(db, "sln_basic_info", nil, nil)
	if err != nil {
		return nil, err
	}

	resp := make([]interface{}, 0)
	for _, value := range dbData {
		el, _ := struct2Map(value.(*SlnBasicInfo), "json")
		resp = append(resp, el)
	}
	return resp, nil
}

func FetchSolutionDetail(db *sql.DB, slnID string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})
	where := map[string]interface{}{
		"sln_no": slnID,
	}

	// 基础表格
	basicData, err := fetchOneData(db, "sln_basic_info", where, nil)
	if err != nil {
		return nil, err
	}
	resp["basic"], _ = struct2Map(basicData, "json")

	// 焊接信息
	weldingData, err := fetchOneData(db, "welding_info", where, nil)
	if err != nil {
		return nil, err
	}
	resp["welding"], _ = struct2Map(weldingData, "json")

	// 方案信息
	weldingSolutionData, err := fetchOneData(db, "welding_devices", where, nil)
	if err == nil {
		resp["sln"], _ = struct2Map(weldingSolutionData, "json")
	} else {
		resp["sln"] = nil
	}

	return resp, nil
}

func CreateSolution(db *sql.DB, params *SolutionParams) (map[string]interface{}, error) {
	return nil, nil
}
