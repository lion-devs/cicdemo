package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-cicd/pkg/metrics"
	"log"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// import other route in this package
	test(router)
	health(router)

	err := router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	if err != nil {
		log.Fatalf(err.Error())
	}

	// set metrics middleware for gin
	monitor := metrics.GetMonitor()
	monitor.Use(router)

	return router
}
