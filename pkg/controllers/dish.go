package pkg

import (
	"net/http"
	pkg "skipthequeue/pkg/models"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
)

type CreateDishInput struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Type         string `json:"type" binding:"required"`
	Availability bool   `json:"availability" binding:"required"`
}

type UpdateDishInput struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Availability bool   `json:"availability"`
}

func CreateDish(c *gin.Context) {
	var input CreateDishInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dish := pkg.Dish{Name: input.Name, Description: input.Description, Type: input.Type, Availability: input.Availability}
	result := utils.DB.Create(&dish)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

func FindAllDish(c *gin.Context) {

}

func FindDishById(c *gin.Context) {

}

func UpdateDish(c *gin.Context) {

}

func DeleteDish(c *gin.Context) {

}
