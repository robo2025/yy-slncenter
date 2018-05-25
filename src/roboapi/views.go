package roboapi

import (
	"github.com/gin-gonic/gin"
	"robodb"
	log "github.com/sirupsen/logrus"
)

// url: /
func (e *GinEnv) viewIndex(c *gin.Context) {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		return
	}

	apiResponse(c, nil, nil)
}

// url: /sln
func (e *GinEnv) viewSolutionList(c *gin.Context) {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		return
	}

	slnList, err := robodb.FetchSolutionList(e.db)
	if err != nil {
		log.Error("Fetch solution list error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, slnList, nil)
	}
}

// post url: /sln
func (e *GinEnv) viewCreateSolution(c *gin.Context) {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		return
	}

	solutionParams := &robodb.SolutionParams{}
	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, nil, err)
		return
	}

	err = robodb.CreateSolution(e.db, solutionParams)
	if err != nil {
		log.Error("Create solution error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, "创建方案成功", nil)
	}
}

// url: /sln/:id
func (e *GinEnv) viewSolutionDetail(c *gin.Context) {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		return
	}

	slnNo := c.Param("id")
	slnDetail, err := robodb.FetchSolutionDetail(e.db, slnNo)
	if err != nil {
		log.Error("Fetch solution detail error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, slnDetail, nil)
	}
}

// put url: /sln/:id
func (e *GinEnv) viewUpdateSolution(c *gin.Context) {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		return
	}

	solutionParams := &robodb.SolutionParams{}
	solutionParams.SlnNo = c.Param("id")

	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, nil, err)
		return
	}

	err = robodb.UpdateSolution(e.db, solutionParams)
	if err != nil {
		log.Error("Update solution error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, "更新方案成功", nil)
	}
}
