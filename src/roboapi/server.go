package roboapi

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"middlewares"
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
	router.Use(middlewares.CorsMiddleware()) // cors

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
	rg.POST("/offer/:id", env.viewOfferSolution) //报价
	rg.PUT("/assign", env.viewAssignSolution)    //指派

	rg.POST("/welding", env.viewCreateWelding) //welding 是焊接方案
	rg.GET("/welding/:id", env.viewDetail)  // 暂时废弃,使用sln/id获取
	rg.PUT("/welding/:id", env.viewUpdateWelding)

	rg.POST("/sewage", env.viewCreateSewage) // sewage 是污水
	rg.GET("/sewage/:id", env.viewDetail)  // 暂时废弃,使用sln/id获取
	rg.PUT("/sewage/:id", env.viewUpdateSewage)

	rg.GET("/sln", env.viewSolutionList) //所有方案列表
	rg.GET("/sln/:id", env.viewDetail) //获取方案详情
	rg.GET("/common/sln/:id", env.viewDetail) //提供接口 获取方案详情

	rg.GET("/log",env.viewGetLog)		 //获取操作记录 log?sln_no=sln_no 获取
	rg.GET("/offer-operation", env.viewGetOfferOperation)      //获取报价操作记录   offer-operation?sln_no=sln_no&sbm_no=sbm_no

	rg.PUT("/checking",env.viewCheckExpire)   // 询价单过期时间检查
}

func registerRPCView(rg *gin.RouterGroup, env *GinEnv) {
	rg.GET("/", env.rpcIndex)
	rg.POST("/sln", env.rpcSolution)
	rg.GET("/sln/:id", env.rpcSolutionDetail)
}
