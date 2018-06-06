package roboapi

import (
	"github.com/gin-gonic/gin"
	"robodb"
	log "github.com/sirupsen/logrus"
)

// url: /
func (e *GinEnv) viewIndex(c *gin.Context) {
	verifyRole := "pass"
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
		return
	}
	apiResponse(c, RespSuccess, nil, "")
}

// url: /sln
func (e *GinEnv) viewSolutionList(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
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
func (e *GinEnv) viewCreateSolution(c *gin.Context) {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
		return
	}

	solutionParams := &robodb.SolutionParams{}
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
func (e *GinEnv) viewSolutionDetail(c *gin.Context) {
	verifyRole := c.Query("role")
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
		return
	}

	slnDetail, err := robodb.FetchSolutionDetail(e.db, c)
	if err != nil {
		log.Error("获取方案细节错误!")
		apiResponse(c, RespNoData, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDetail, "")
	}
}

// put url: /sln/:id
func (e *GinEnv) viewUpdateSolution(c *gin.Context) {
	verifyRole := "customer"
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
		return
	}

	solutionParams := &robodb.SolutionParams{}
	solutionParams.SlnNo = c.Param("id")

	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	err = robodb.UpdateSolution(e.db, solutionParams, c)
	if err != nil {
		log.Error("更新方案列表错误!")
		apiResponse(c, RespFailed, nil, err.Error())
	} else {
		apiResponse(c, RespSuccess, nil, "更新方案成功")
	}
}

// post url: /offer/:id
func (e *GinEnv) viewOfferSolution(c *gin.Context) {
	verifyRole := "supplier"
	if err := checkAuthRole(c, verifyRole); err != nil {
		apiResponse(c, RespTokenFailed, nil, err.Error())
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
