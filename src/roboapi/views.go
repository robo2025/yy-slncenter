package roboapi

import (
	"github.com/gin-gonic/gin"
	"robodb"
	log "github.com/sirupsen/logrus"
)

// url: /
func (e *GinEnv) viewIndex(c *gin.Context) {
	apiResponse(c, nil, nil)
}

// url: /sln
func (e *GinEnv) viewSolutionList(c *gin.Context) {
	slnList, err := robodb.FetchSolutionList(e.db)
	if err != nil {
		log.Error("Fetch solution list error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, slnList, nil)
	}
}

// post url: /sln
func (e *GinEnv) createSolution(c *gin.Context) {
	solutionParams := &robodb.SolutionParams{}
	err := c.BindJSON(solutionParams)
	if err != nil {
		apiResponse(c, nil, err)
		return
	}

	//resp, err := robodb.CreateSolution(e.db, solutionParams)

	if err != nil {
		log.Error("create solution error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, solutionParams, nil)
	}
}

// url: /sln/:id
func (e *GinEnv) viewSolutionDetail(c *gin.Context) {
	slnNo := c.Param("id")
	slnDetail, err := robodb.FetchSolutionDetail(e.db, slnNo)
	if err != nil {
		log.Error("Fetch solution detail error!")
		apiResponse(c, nil, err)
	} else {
		apiResponse(c, slnDetail, nil)
	}
}
