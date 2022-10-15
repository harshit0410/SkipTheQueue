package controllers

import (
	"errors"
	"net/http"

	"skipthequeue/pkg/models"
	"skipthequeue/pkg/services"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateDishInput struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Type         string `json:"type" binding:"required"`
	Availability bool   `json:"availability" binding:"required"`
	Outlet       int    `json:"outlet" binding:"required"`
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

	dish := models.Dish{Name: input.Name, Description: input.Description, Type: input.Type, Availability: input.Availability, Outlet: uint(input.Outlet)}
	result := utils.DB.Create(&dish)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

func FindAllDish(c *gin.Context) {
	var dishes []models.Dish

	outletId := c.Param("outletId")

	_, err := services.FindOutletByIdService(outletId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := utils.DB.Where("outlet = ?", outletId).Find(&dishes)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dishes})
}

func FindDishById(c *gin.Context) {
	var dish models.Dish

	outletId := c.Param("outletId")
	dishId := c.Param("dishId")

	_, err := services.FindOutletByIdService(outletId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := utils.DB.Where("outlet = ?", outletId).First(&dish, dishId)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

func UpdateDish(c *gin.Context) {
	var dish models.Dish
	outletId := c.Param("ouletId")
	dishId := c.Param("dishId")

	_, err := services.FindOutletByIdService(outletId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := utils.DB.Where("outlet = ?", outletId).First(&dish, dishId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	var updatedDishInput UpdateDishInput

	if err := c.ShouldBindJSON(&updatedDishInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedResult := utils.DB.Model(&dish).Updates(updatedDishInput)

	if updatedResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": updatedResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

func DeleteDish(c *gin.Context) {
	var dish models.Dish

	outletId := c.Param("outletId")
	dishId := c.Param("DishId")

	_, err := services.FindOutletByIdService(outletId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := utils.DB.Where("outlet = ?", outletId).First(&dish, dishId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	deletedResult := utils.DB.Delete(&dish, dishId)

	if deletedResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": deletedResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}
