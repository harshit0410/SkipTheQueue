package controllers

import (
	"net/http"
	"skipthequeue/pkg/models"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
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
		Type:     "takwAway",
	}
	err = utils.DB.Create(order).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	orderDetails, finalAmount, err := createOrderDetails(orderInput.OrderDetails, order.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = utils.DB.Create(orderDetails).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	order.FinalAmount = finalAmount
	err = utils.DB.Save(order).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
