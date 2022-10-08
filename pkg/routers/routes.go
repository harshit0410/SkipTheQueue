package pkg

import (
	pkg "skipthequeue/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", pkg.Ping)

	router.GET("/outlets", pkg.FindAllOutlet)
	router.POST("/outlets", pkg.CreateOutlet)
	router.GET("/outlets/:id", pkg.FindOutletById)
	router.PATCH("/outlets/:id", pkg.UpdateOutlet)
	router.DELETE("/outlets/:id", pkg.DeleteOutlet)
}
