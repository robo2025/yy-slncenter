package roboapi

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type GinEnv struct {
	db *gorm.DB
}

var DeployMode string

func StartWebService(bindAddr string, db *gorm.DB) {

	DeployMode = os.Getenv("ROBO_DEPLOY_MODE")
	// Set gin context
	if DeployMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()
	env := &GinEnv{db: db}

	// Set API
	apiGroup := router.Group("/v1")
	apiGroup.Use(cors.Default())
	apiGroup.Use(SSOAuth())
	registerApiView(apiGroup, env)

	// Set RPC
	rpcGroup := router.Group("/rpc")
	rpcGroup.Use(cors.Default())
	registerRPCView(rpcGroup, env)

	// Run server
	fmt.Println("API listen on: ", bindAddr)
	router.Run(bindAddr)
}

func registerApiView(rg *gin.RouterGroup, env *GinEnv) {
	rg.GET("/", env.viewIndex)
	rg.GET("/sln", env.viewSolutionList)
	rg.POST("/sln", env.viewCreateSolution)
	rg.GET("/sln/:id", env.viewSolutionDetail)
	rg.PUT("/sln/:id", env.viewUpdateSolution)
	rg.POST("/offer/:id", env.viewOfferSolution)
}

func registerRPCView(rg *gin.RouterGroup, env *GinEnv) {
	rg.GET("/", env.rpcIndex)
	rg.POST("/sln", env.rpcSolution)
	rg.GET("/sln/:id", env.rpcSolutionDetail)
}
