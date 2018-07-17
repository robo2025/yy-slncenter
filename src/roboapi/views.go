package roboapi

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"robodb"
)

// url: /
func (e *GinEnv) viewIndex(c *gin.Context) {
	verifyRole := "pass"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}
	apiResponse(c, RespSuccess, nil, "")
}

// url: /sln
func (e *GinEnv) viewSolutionList(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	slnList, err := robodb.FetchSolutionList(e.db, c)
	if err != nil {
		log.Error("获取方案列表错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnList, "")
	}
}

// post url: /sln
func (e *GinEnv) viewCreateWelding(c *gin.Context) {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	solutionParams := &robodb.WeldingParams{}
	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.CreateSolution(e.db, solutionParams, c)
	if err != nil {
		log.Error("创建方案错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "创建方案成功")
	}
}

// url: /sln/:id
func (e *GinEnv) viewWeldingDetail(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}
	slnID := c.Param("id")
	basicInfo := &robodb.SlnBasicInfo{}
	e.db.Where("sln_no = ?", slnID).First(basicInfo)
	if basicInfo.SlnType != "welding" {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespFailed, nil, "获取方案细节错误!该方案id不是welding类")
		return
	}


	slnDetail, err := robodb.FetchWeldingDetail(e.db, c)
	if err != nil {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDetail, "")
	}
}

// put url: /sln/:id
func (e *GinEnv) viewUpdateWelding(c *gin.Context) {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	solutionParams := &robodb.WeldingParams{}
	solutionParams.SlnNo = c.Param("id")

	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.UpdateWelding(e.db, solutionParams, c)
	if err != nil {
		log.Error("更新方案列表错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "更新方案成功")
	}
}

// post url: /offer/:id
func (e *GinEnv) 	viewOfferSolution(c *gin.Context) {
	verifyRole := "supplier"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}
	// 验证
	slnID := c.Param("id")
	assignInfo := &robodb.SlnAssign{}
	uid := c.MustGet("uid").(int)
	e.db.Where("sln_no = ?", slnID).First(assignInfo)
	if assignInfo.SupplierId != uid {
		log.Error("请求报价错误!")
		apiResponse(c, RespFailed, nil, "请求报价错误!该方案并没有指派给这个供应商")
		return
	}
	supplierInfo := &robodb.SlnSupplierInfo{}

	e.db.Where("sln_no = ?", slnID).First(supplierInfo)
	if supplierInfo.SlnNo == slnID && supplierInfo.UserID == uid && supplierInfo.TotalPrice != 0 {
		log.Error("请求报价错误!")
		apiResponse(c, RespFailed, nil, "不能重复报价")
		return
	}

	// 解析请求
	offerParams := &robodb.OfferParams{}
	offerParams.SlnNo = c.Param("id")
	err := c.BindJSON(offerParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.OfferSolution(e.db, offerParams, c)
	if err != nil {
		log.Error("方案报价错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "")
	}
}

// to add sewage viewCreateSewage

// post url: /sewage
func (e *GinEnv) viewCreateSewage(c *gin.Context) {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	solutionParams := &robodb.SewageParams{}
	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.CreateSewage(e.db, solutionParams, c)
	if err != nil {
		log.Error("创建方案错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "创建方案成功")
	}
}

// url: /sln/:id
func (e *GinEnv) viewSewageDetail(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	slnID := c.Param("id")
	basicInfo := &robodb.SlnBasicInfo{}
	e.db.Where("sln_no = ?", slnID).First(basicInfo)
	if basicInfo.SlnType != "sewage" {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespFailed, nil, "获取方案细节错误!该方案id不是sewage类")
		return
	}


	slnDetail, err := robodb.FetchSewageDetail(e.db, c)
	if err != nil {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDetail, "")
	}
}

func (e *GinEnv) viewUpdateSewage(c *gin.Context)  {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	sewageParams := &robodb.SewageParams{}
	sewageParams.SlnNo = c.Param("id")

	err := c.BindJSON(sewageParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.UpdateSewage(e.db, sewageParams, c)
	if err != nil {
		log.Error("更新方案列表错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "更新方案成功")
	}
}

// url: /sln/:id
func (e *GinEnv) viewDetail(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	slnDetail, err := robodb.FetchDetail(e.db, c)
	if err != nil {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDetail, "")
	}
}

// url: /assign/:id
func (e *GinEnv) viewAssignSolution(c *gin.Context)  {
	verifyRole := "admin"
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	// 解析请求
	assignParams := &robodb.AssignParams{}
	err := c.BindJSON(assignParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.AssignSolution(e.db, assignParams, c)
	if err != nil {
		log.Error("方案指派错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "")
	}
}

// url: /log/:id
func (e *GinEnv) viewGetLog(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		return
	}

	slnDetail, err := robodb.FetchLog(e.db, c)
	if err != nil {
		log.Error("获取操作记录细节错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDetail, "")
	}
}