package roboapi

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type GinEnv struct {
	db *gorm.DB
}

func StartWebService(bindAddr string, db *gorm.DB) {

	gin.SetMode(gin.ReleaseMode)
	env := &GinEnv{db: db}

	router := gin.Default()
	// use cors
	router.Use(cors.Default())
	// use jwt
	router.Use(JWTAuth())

	// Register router
	routerGroup := router.Group("/v1")
	registerApiView(routerGroup, env)

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
