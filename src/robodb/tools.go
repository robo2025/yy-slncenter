package robodb

import (
	"strings"
	"sort"
	"reflect"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// 字符串数组去重复
func filterDuplicateString(dbList []string) []string {
	if len(dbList) < 2 {
		return dbList
	}

	dbMap := make(map[string]string)
	for _, el := range dbList {
		trimStr := strings.TrimSpace(el)
		if trimStr != "" {
			dbMap[trimStr] = trimStr
		}
	}

	dbResp := make([]string, 0)
	for key := range dbMap {
		dbResp = append(dbResp, key)
	}

	sort.Strings(dbResp)
	return dbResp
}

// 数组去重复
func filterDuplicate(dbList []interface{}) []interface{} {
	if len(dbList) < 2 {
		return dbList
	}

	dbResp := make([]interface{}, 0)

	// 判断传入的值
	v := reflect.ValueOf(dbList[0])
	switch v.Kind() {
	case reflect.String:
		dbMap := make(map[string]interface{})
		for _, el := range dbList {
			trimStr := strings.TrimSpace(el.(string))
			if trimStr != "" {
				dbMap[trimStr] = trimStr
			}
		}
		for key := range dbMap {
			dbResp = append(dbResp, key)
		}
	case reflect.Int:
		dbMap := make(map[int]interface{})
		for _, el := range dbList {
			dbMap[el.(int)] = el
		}
		for key := range dbMap {
			dbResp = append(dbResp, key)
		}
	case reflect.Float64:
		dbMap := make(map[float64]interface{})
		for _, el := range dbList {
			dbMap[el.(float64)] = el
		}
		for key := range dbMap {
			dbResp = append(dbResp, key)
		}
	default:
		log.Debug("Not supported filter type: ", v.Kind())
	}
	return dbResp
}

// Structure 转 Map
// Reference https://stackoverflow.com/questions/23589564/function-for-converting-a-struct-to-map-in-golang/25117810#25117810
func struct2Map(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// only accept struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a structure field
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil
}
