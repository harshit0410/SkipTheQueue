package controllers

import (
	"errors"
	"net/http"
	"skipthequeue/core/models"
	"skipthequeue/core/services"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateOrderDetailInput struct {
	DishId   uint `json:"dishId" binding:"required"`
	Quantity uint `json:"quantity" binding:"required"`
}

type CreateOrderInput struct {
	OutletId     uint                     `json:"outletId" binding:"required"`
	OrderDetails []CreateOrderDetailInput `json:"orderDetails" binding:"required"`
}

func createOrderDetails(orderDetailsInput []CreateOrderDetailInput, orderId uint) ([]models.OrderDetails, float32, error) {
	var dishes []models.Dish
	var ids []uint
	var orderDetails []models.OrderDetails
	var totalAmount float32 = 0

	for _, element := range orderDetailsInput {
		ids = append(ids, element.DishId)
	}

	result := utils.DB.Select("Id", "Price").Where("id in ?", ids).Find(&dishes)
	if result.Error != nil {
		return orderDetails, float32(totalAmount), result.Error
	}

	dishPriceMap := make(map[uint]float32)

	for _, dish := range dishes {
		dishPriceMap[dish.Id] = dish.Price
	}

	for _, element := range orderDetailsInput {
		totalPrice := dishPriceMap[element.DishId] * float32(element.Quantity)
		orderDetails = append(orderDetails, models.OrderDetails{
			OrderId:    orderId,
			DishId:     element.DishId,
			Quantity:   element.Quantity,
			Price:      dishPriceMap[element.DishId],
			TotalPrice: totalPrice,
		})

		totalAmount += totalPrice
	}

	return orderDetails, float32(totalAmount), nil

}

func CreateOrder(c *gin.Context) {
	var orderInput CreateOrderInput
	var err error
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		OutletId: orderInput.OutletId,
		Status:   "ordered",
		Type:     "takeAway",
	}
	err = utils.DB.Create(&order).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	orderDetails, finalAmount, err := createOrderDetails(orderInput.OrderDetails, order.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = utils.DB.Create(&orderDetails).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	order.FinalAmount = finalAmount
	err = utils.DB.Save(&order).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

type UpdateOrderStatusInput struct {
	Status string `json:"status" binding:"required"`
}

func UpdateOrderStatus(c *gin.Context) {
	var updateOrderStatusInput models.Order
	orderId := c.Param("orderId")

	order, err := services.FindOrderByIdService(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&updateOrderStatusInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedResult := utils.DB.Model(&order).Updates(updateOrderStatusInput)

	if updatedResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": updatedResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func FindOrderById(c *gin.Context) {
	var order models.Order

	id := c.Param("id")

	result := utils.DB.First(&order, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recond not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
