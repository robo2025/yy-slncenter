package robodb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

func InitDB(sqlURL string) (*sql.DB, error) {

	db, err := sql.Open("mysql", sqlURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 数据库 API
func fetchOneData(db *sql.DB, table string, where map[string]interface{}, field []string) (interface{}, error) {
	if db == nil {
		return nil, errors.New("sql.DB object couldn't be nil")
	}

	cond, vals, err := builder.BuildSelect(table, where, field)
	if err != nil {
		return nil, err
	}

	row, err := db.Query(cond, vals...)
	if err != nil || row == nil {
		return nil, err
	}
	defer row.Close()

	scanner.SetTagName("json")
	var resp interface{}

	switch table {
	case "sln_basic_info":
		var scanData *SlnBasicInfo
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}
		resp = scanData

	case "welding_info":
		var scanData *WeldingInfo
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}
		resp = scanData

	case "welding_devices":
		var scanData *WeldingDevices
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}
		resp = scanData

	default:
		return nil, errors.New("table name error")
	}

	return resp, nil
}

// 数据库 API
func fetchMultiData(db *sql.DB, table string, where map[string]interface{}, field []string) ([]interface{}, error) {

	if db == nil {
		return nil, errors.New("sql.DB object couldn't be nil")
	}

	cond, vals, err := builder.BuildSelect(table, where, field)
	if err != nil {
		return nil, err
	}

	row, err := db.Query(cond, vals...)
	if err != nil || row == nil {
		return nil, err
	}
	defer row.Close()

	scanner.SetTagName("json")
	var resp []interface{}

	switch table {
	case "sln_basic_info":
		var scanData []*SlnBasicInfo
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}

		resp = make([]interface{}, len(scanData))
		for k, v := range scanData {
			resp[k] = v
		}

	case "welding_info":
		var scanData []*WeldingInfo
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}

		resp = make([]interface{}, len(scanData))
		for k, v := range scanData {
			resp[k] = v
		}

	case "welding_devices":
		var scanData []*WeldingDevices
		err = scanner.Scan(row, &scanData)
		if err != nil {
			return nil, err
		}

		resp = make([]interface{}, len(scanData))
		for k, v := range scanData {
			resp[k] = v
		}

	default:
		return nil, errors.New("table name error")
	}

	return resp, nil
}
