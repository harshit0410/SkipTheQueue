package routes

import (
	"skipthequeue/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)

	router.GET("/outlets", controllers.FindAllOutlet)
	router.POST("/outlets", controllers.CreateOutlet)
	router.GET("/outlets/:id", controllers.FindOutletById)
	router.PATCH("/outlets/:id", controllers.UpdateOutlet)
	router.DELETE("/outlets/:id", controllers.DeleteOutlet)
}
