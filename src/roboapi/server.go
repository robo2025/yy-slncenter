package roboapi

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"database/sql"
	"fmt"
)

type GinEnv struct {
	db *sql.DB
}

func StartWebService(bindAddr string, db *sql.DB) {

	gin.SetMode(gin.ReleaseMode)
	env := &GinEnv{db: db}

	router := gin.Default()
	router.Use(cors.Default())

	// Register router
	routerGroup := router.Group("/v1")
	registerApiView(routerGroup, env)

	fmt.Println("API listen on: ", bindAddr)
	router.Run(bindAddr)
}

func registerApiView(rg *gin.RouterGroup, env *GinEnv) {
	rg.GET("/", env.viewIndex)
	rg.GET("/sln", env.viewSolutionList)
	rg.POST("/sln", env.createSolution)
	rg.GET("/sln/:id", env.viewSolutionDetail)
}
