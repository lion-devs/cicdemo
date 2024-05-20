// Package router provides HTTP routing functionality.
//
// swagger:meta
package router

import (
	"github.com/JustinHung0407/cicdemo/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAll godoc
// @summary Test the alert service and OPA service.
// @description Returns the status of the alert service and OPA service.
// @tags test
// @produce json
// @success 200 {object} Response
// @router /alert_test [get]
// @response default {object} Response
func test(router *gin.Engine) {

	// Setup Response
	response := model.Response{Message: "ok"}
	badGateway := model.Response{Status: 503, Message: "Service Unavailable Test"}

	test := router.Group("/")
	{
		test.GET("/alert_test", func(c *gin.Context) {
			c.Status(http.StatusServiceUnavailable)
			c.JSON(http.StatusServiceUnavailable, badGateway)
		})

		test.GET("/opa_test", func(c *gin.Context) {
			c.JSON(http.StatusOK, response)
		})
	}
}
