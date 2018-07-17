package robodb

import (
	"github.com/jinzhu/gorm"
	"errors"
	"time"
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


