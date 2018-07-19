package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
	"time"
	"roboutil"
	"fmt"
	"strconv"
	"strings"
)

// 准备指派方案数据
func prepareAssignData(params *AssignParams) *AssignParams {
	currentTime := time.Now().Unix()
	params.SlnAssign.AddTime = int(currentTime)
	return params
}

// 写入指派方案数据
func writeAssignData(db *gorm.DB, params *AssignParams) error {
	var err error

	slnBasicInfo := &SlnBasicInfo{}
	db.Where("sln_no = ?", params.SlnAssign.SlnNo).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return errors.New("找不到相应方案")
	}

	if slnBasicInfo.SlnStatus != string(SlnStatusPublish) {
		return errors.New("该方案不可以指派")
	}
	// 已发布状态下的未指派订单才可以指派
	if slnBasicInfo.AssignStatus == string(AssignStatusY) {
		return errors.New("该方案不可以指派")
	}
	// 指派后让运营看到指派给了哪位供应商
	SlnSupplierInfo := &SlnSupplierInfo{}
	SlnSupplierInfo.SlnNo = params.SlnAssign.SlnNo
	SlnSupplierInfo.UserID = params.SlnAssign.SupplierId

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

	// 创建sln_supplier_info表信息
	err = tx.Create(SlnSupplierInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 更新 sln_basic_info 表
	tx.Model(slnBasicInfo).Updates(SlnBasicInfo{
		AssignStatus: string(AssignStatusY),
		SupplierID: params.SlnAssign.SupplierId,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	// 写入 sln_assign 表
	err = tx.Create(params.SlnAssign).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}


// 查询供应商报价细节
func readOfferData(db *gorm.DB, slnID string, uid int) (*OfferParams, error) {
	slnSupplierInfo := &SlnSupplierInfo{}
	slnDevice := []SlnDevice{}
	weldingTechParams := []WeldingTechParam{}
	slnSupport := []SlnSupport{}
	slnFile := []SlnFile{}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).First(slnSupplierInfo)
	if slnSupplierInfo.ID == 0 {
		return nil, errors.New("找不到相应报价")
	}

	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&weldingTechParams)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnSupport)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnDevice)
	db.Where("sln_no = ? AND user_id = ?", slnID, uid).Find(&slnFile)

	resp := &OfferParams{
		SlnNo:            slnID,
		SlnSupplierInfo:  slnSupplierInfo,
		SlnDevice:        slnDevice,
		WeldingTechParam: weldingTechParams,
		SlnSupport:       slnSupport,
		SlnFile:          slnFile,
	}

	return resp, nil
}

// 准备方案报价数据
func prepareOfferData(params *OfferParams, uid int) *OfferParams {
	// 准备数据
	slnNo := strings.TrimSpace(params.SlnNo)

	params.UID = uid

	// sln_supplier_info
	slnSupplierInfo := params.SlnSupplierInfo
	currentDate := time.Now()
	if slnSupplierInfo != nil {
		slnSupplierInfo.SlnNo = slnNo
		slnSupplierInfo.UserID = uid
		slnSupplierInfo.ExpiredDate = int(currentDate.AddDate(0, 0, 30).Unix())

	}

	// welding_device 表
	slnDevice := make([]SlnDevice, 0)
	if len(params.SlnDevice) != 0 {
		for _, el := range params.SlnDevice {
			el.SlnNo = slnNo
			el.UserID = uid
			el.SlnRole = "S"
			slnDevice = append(slnDevice, el)
		}
	}

	// sln_support 表
	slnSupport := make([]SlnSupport, 0)
	if len(params.SlnSupport) != 0 {
		for _, el := range params.SlnSupport {
			el.SlnNo = slnNo
			el.UserID = uid
			slnSupport = append(slnSupport, el)
		}
	}

	// welding_tech_param 表
	weldingTechParam := make([]WeldingTechParam, 0)
	if len(params.WeldingTechParam) != 0 {
		for _, el := range params.WeldingTechParam {
			el.SlnNo = slnNo
			el.UserID = uid
			weldingTechParam = append(weldingTechParam, el)
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
	resp := &OfferParams{
		SlnNo:            slnNo,
		UID:              uid,
		SlnSupplierInfo:  slnSupplierInfo,
		SlnDevice:        slnDevice,
		SlnSupport:       slnSupport,
		WeldingTechParam: weldingTechParam,
		SlnFile:          slnFile,
	}
	return resp
}

// 写入报价方案数据   todo 创建时,指派状态默认为 未指派 (W)   Y
func writeOfferData(db *gorm.DB, params *OfferParams) error {
	var err error
	var sbmNo int

	slnBasicInfo := &SlnBasicInfo{}
	db.Where("sln_no = ?", params.SlnNo).First(slnBasicInfo)
	if slnBasicInfo.ID == 0 {
		return errors.New("找不到相应方案")
	}

	if slnBasicInfo.SlnStatus != string(SlnStatusPublish) {
		return errors.New("该方案不可以指派")
	}
	slnSupplierInfo := &SlnSupplierInfo{}
	db.Where("sln_no = ?", params.SlnNo).First(slnSupplierInfo)
	// 准备操作日志信息
	currentDate := time.Now()
	userName := roboutil.HttpGet(params.SlnSupplierInfo.UserID)
	Operator := fmt.Sprintf("供应商(%s)", userName)
	price := strconv.FormatFloat(params.SlnSupplierInfo.TotalPrice, 'f', 2, 64)
	content := fmt.Sprintf("该方案询价单报价%s元", price)
	operationLog := &OperationLog{}
	operationLog.SlnNo = params.SlnNo
	operationLog.OperationType = "报价"
	operationLog.Operator = Operator
	operationLog.Content = content
	operationLog.AddTime = int(currentDate.Unix())
	operationLog.OperatorId = params.SlnSupplierInfo.UserID

	//准备操作记录信息(device的更改)
	db.Where("sln_no = ?", params.SlnNo).Find(&[]SlnSupplierInfo{}).Count(&sbmNo)
	oldSlnDevice := []SlnDevice{}
	offerOperation := []OfferOperation{}
	db.Where("sln_no = ? AND user_id = ?", params.SlnNo, slnBasicInfo.CustomerID).Find(&oldSlnDevice)
	newSlnDevice := params.SlnDevice
	fmt.Println(oldSlnDevice, newSlnDevice)
	var oldMap map[string]SlnDevice
	oldMap = make(map[string]SlnDevice)
	for i := 0; i < len(oldSlnDevice); i++ {
		oldMap[oldSlnDevice[i].DeviceID] = oldSlnDevice[i]
	}
	for i := 0; i < len(newSlnDevice); i++ {
		_, ok := oldMap[newSlnDevice[i].DeviceID]
		if !ok {
			price := strconv.FormatFloat(newSlnDevice[i].DevicePrice, 'f', 2, 64)
			operation := OfferOperation{
				SlnNo:         params.SlnNo,
				SbmNo:         sbmNo,
				Role:          int(SupplierUser),
				OperatingPart: newSlnDevice[i].DeviceType,
				OperatingType: "添加",
				Content: fmt.Sprintf("添加 组成部分:%s;产品名称:%s;型号: %s;品牌: %s;单价: %s;数量: %d;", newSlnDevice[i].DeviceComponent, newSlnDevice[i].DeviceName, newSlnDevice[i].DeviceModel, newSlnDevice[i].BrandName, price, newSlnDevice[i].DeviceNum),
				AddTime: int(time.Now().Unix()),
			}
			offerOperation = append(offerOperation, operation)
		}
	}
	var newMap map[string]SlnDevice
	newMap = make(map[string]SlnDevice)
	for i := 0; i < len(newSlnDevice); i++ {
		if newSlnDevice[i].DeviceID != "" {
			newMap[newSlnDevice[i].DeviceID] = newSlnDevice[i]
		} else {
			price := strconv.FormatFloat(newSlnDevice[i].DevicePrice, 'f', 2, 64)
			operation := OfferOperation{
				SlnNo:         params.SlnNo,
				SbmNo:         sbmNo,
				Role:          int(SupplierUser),
				OperatingPart: newSlnDevice[i].DeviceType,
				OperatingType: "添加",
				Content: fmt.Sprintf("添加 组成部分：%s;产品名称：%s;型号: %s;品牌: %s;单价:%s;数量: %d", newSlnDevice[i].DeviceComponent, newSlnDevice[i].DeviceName,
					newSlnDevice[i].DeviceModel, newSlnDevice[i].BrandName, price, newSlnDevice[i].DeviceNum),
				AddTime: int(time.Now().Unix()),
			}
			offerOperation = append(offerOperation, operation)
		}
	}
	constantDevice := []SlnDevice{} // 添加未删除的原询价单设备,检查有没有修改属性
	for i := 0; i < len(oldSlnDevice); i++ {
		v, ok := newMap[oldSlnDevice[i].DeviceID]
		if !ok {
			price := strconv.FormatFloat(oldSlnDevice[i].DevicePrice, 'f', 2, 64)
			operation := OfferOperation{
				SlnNo:         params.SlnNo,
				SbmNo:         sbmNo,
				Role:          int(SupplierUser),
				OperatingPart: oldSlnDevice[i].DeviceType,
				OperatingType: "删除",
				Content: fmt.Sprintf("删除 组成部分：%s;产品名称：%s;型号: %s;品牌: %s;单价: %s;数量: %d;", oldSlnDevice[i].DeviceComponent, oldSlnDevice[i].DeviceName,
					oldSlnDevice[i].DeviceModel, oldSlnDevice[i].BrandName, price, oldSlnDevice[i].DeviceNum),
				AddTime: int(time.Now().Unix()),
			}
			offerOperation = append(offerOperation, operation)
		} else {
			constantDevice = append(constantDevice, v)
		}
	}
	//比较没有添加和删除的设备,属性价格和数量有无变化
	for i:=0;i<len(constantDevice);i++ {
		for k:=0;k<len(oldSlnDevice);k++{
			if constantDevice[i].DeviceID == oldSlnDevice[k].DeviceID{
				if constantDevice[i].DeviceNum !=oldSlnDevice[k].DeviceNum || constantDevice[i].DevicePrice !=oldSlnDevice[k].DevicePrice{
					price := strconv.FormatFloat(constantDevice[k].DevicePrice, 'f', 2, 64)
					operation := OfferOperation{
						SlnNo:         params.SlnNo,
						SbmNo:         sbmNo,
						Role:          int(SupplierUser),
						OperatingPart: oldSlnDevice[i].DeviceType,
						OperatingType: "修改",
						Content: fmt.Sprintf("修改 组成部分：%s;产品名称：%s;型号: %s;品牌: %s;单价: %s;数量: %d;", oldSlnDevice[i].DeviceComponent, oldSlnDevice[i].DeviceName,
							oldSlnDevice[i].DeviceModel, oldSlnDevice[i].BrandName, price, constantDevice[k].DeviceNum),
						AddTime: int(time.Now().Unix()),
					}
					offerOperation = append(offerOperation, operation)
				}
			}
		}
	}
	// prepare 技术支持
	//db.Where("sln_no = ?", params.SlnNo).Find(&[]SlnSupplierInfo{}).Count(&sbmNo)
	for i := 0; i < len(params.SlnSupport); i++ {
		fmt.Println(params.SlnSupport[i].Name)
		price := strconv.FormatFloat(params.SlnSupport[i].Price, 'f', 1, 64)
		operation := OfferOperation{
			SlnNo:         params.SlnNo,
			SbmNo:         sbmNo,
			Role:          int(SupplierUser),
			OperatingPart: "技术支持",
			OperatingType: "选择",
			Content:       fmt.Sprintf("选择 项目名称：%s；输入价格：%s", params.SlnSupport[i].Name, price),
			AddTime:       int(time.Now().Unix()),
		}
		offerOperation = append(offerOperation, operation)
	}

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
	if slnBasicInfo.SupplierID == params.SlnSupplierInfo.UserID {
		tx.Model(slnBasicInfo).Updates(SlnBasicInfo{
			SlnStatus:     string(SlnStatusOffer),
			SupplierID:    params.SlnSupplierInfo.UserID,
			SupplierPrice: params.SlnSupplierInfo.TotalPrice,
			AssignStatus:  string(AssignStatusY),
			SupplierName:  roboutil.HttpGet(params.SlnSupplierInfo.UserID),
			SpDate:        int(time.Now().Unix()), //报价日期
		})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 写入 sln_supplier_info 表
	tx.Model(slnSupplierInfo).Updates(SlnSupplierInfo{
		TotalPrice:   params.SlnSupplierInfo.TotalPrice,
		FreightPrice: params.SlnSupplierInfo.FreightPrice,
		PayRatio:     params.SlnSupplierInfo.PayRatio,
		ExpiredDate:  params.SlnSupplierInfo.ExpiredDate,
		DeliveryDate: params.SlnSupplierInfo.DeliveryDate,
		SlnDesc:      params.SlnSupplierInfo.SlnDesc,
		SlnNote:      params.SlnSupplierInfo.SlnNote,
	})
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

	// 写入 sln_support 表
	if len(params.SlnSupport) != 0 {
		for _, el := range params.SlnSupport {
			err = tx.Create(&el).Error

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_tech_param 表
	if len(params.WeldingTechParam) != 0 {
		for _, el := range params.WeldingTechParam {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 写入 welding_file 表
	if len(params.SlnFile) != 0 {
		for _, el := range params.SlnFile {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	// 写入operation_log表
	err = tx.Create(operationLog).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 写入offer_operation表
	if len(offerOperation) != 0 {
		for _, el := range offerOperation {
			err = tx.Create(&el).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}