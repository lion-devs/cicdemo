package router

import (
	"go-cicd/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAll godoc
// @summary Checks the health status of the service.
// @description Returns the health status of the service.
// @tags health
// @produce json
// @success 200 {object} Response
// @router / [get]
// @response default {object} Response
func health(router *gin.Engine) {

	// Setup Response
	response := &model.Response{
		Status:  0,
		Message: "ok",
	}

	health := router.Group("/")
	{
		health.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, response)
		})

		health.GET("/__health", func(c *gin.Context) {
			c.JSON(http.StatusOK, response)
		})

		health.GET("/proto", func(c *gin.Context) {
			c.ProtoBuf(http.StatusOK, model.NewResponseMessage(0, "ok"))
		})
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
