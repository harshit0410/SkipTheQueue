package controllers

import (
	"errors"
	"net/http"

	"skipthequeue/core/models"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type CreateOutletInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	PinCode string `json:"pincode" binding:"required"`
	City    string `json:"city" binding:"required"`
}

type UpdateOutletInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	PinCode string `json:"pincode"`
	City    string `json:"city"`
}

func CreateOutlet(c *gin.Context) {
	var input CreateOutletInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outlet := models.Outlet{Name: input.Name, Address: input.Address, City: input.City, PinCode: input.PinCode}
	result := utils.DB.Create(&outlet)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outlet})
}

func FindAllOutlet(c *gin.Context) {
	var outlets []models.Outlet

	result := utils.DB.Find(&outlets)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outlets})
}

func FindOutletById(c *gin.Context) {
	var outlet models.Outlet

	id := c.Param("id")

	result := utils.DB.First(&outlet, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outlet})
}

func UpdateOutlet(c *gin.Context) {
	var outlet models.Outlet
	id := c.Param("id")

	result := utils.DB.First(&outlet, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	var updatedOutletInput models.Outlet

	if err := c.ShouldBindJSON(&updatedOutletInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mergo.Merge(&outlet, updatedOutletInput, mergo.WithOverride); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedResult := utils.DB.Model(&outlet).Updates(outlet)

	if updatedResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": updatedResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outlet})
}

func DeleteOutlet(c *gin.Context) {
	var outlet models.Outlet

	id := c.Param("id")
	result := utils.DB.First(&outlet, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	deletedResult := utils.DB.Delete(&outlet, id)

	if deletedResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": deletedResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outlet})
}
