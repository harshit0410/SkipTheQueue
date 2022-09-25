package pkg

import (
	pkg "skipthequeue/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/outlets", pkg.FindAllOutlet)
	router.POST("/outlets", pkg.CreateOutlet)
}
