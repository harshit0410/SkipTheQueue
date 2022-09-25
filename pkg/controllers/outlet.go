package pkg

import (
	"net/http"
	pkg "skipthequeue/pkg/models"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
)

type CreateOutletInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	PinCode string `json:"pincode" binding:"required"`
	City    string `json:"city" binding:"required"`
}

func CreateOutlet(c *gin.Context) {
	var input CreateOutletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outlet := pkg.Outlet{Name: input.Name, Address: input.Address, City: input.City, PinCode: input.PinCode}
	utils.DB.Create(&outlet)
	c.JSON(http.StatusOK, gin.H{"data": outlet})
}

func FindAllOutlet(c *gin.Context) {
	var outlets []pkg.Outlet

	utils.DB.Find(&outlets)

	c.JSON(http.StatusOK, gin.H{"data": outlets})
}

func FindOutletById(c *gin.Context) {

}

func UpdateOutlet(c *gin.Context) {

}

func DeleteOutlet(c *gin.Context) {

}
