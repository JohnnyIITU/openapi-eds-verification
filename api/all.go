package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openapi-eds-verification/api/response"
	"openapi-eds-verification/helpers"
)

func AllApi(r *gin.Engine) {
	group := r.Group("v1")

	group.GET("get-info", func(c *gin.Context) {
		x509 := c.GetString("x509")

		if x509 == "" {
			c.AbortWithStatusJSON(400, response.NewErrorResponse("x509 is empty"))
		}

		params := map[string]string{
			"x509": x509,
		}

		resp, err := helpers.NcaNodeRequest("X509.info", params)

		if err != nil {
			c.AbortWithStatusJSON(resp.StatusCode, err.Error())
		}
	})

	group.POST("validate-certificate", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"work": 1,
		})
	})

	group.POST("validate-certificate", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"work": 1,
		})
	})
}
