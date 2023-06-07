package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Openings",
			})
		})
	}
}
