package services

import (
	"errors"
	"skipthequeue/core/models"

	"skipthequeue/utils"

	"gorm.io/gorm"
)

func FindOrderByIdService(orderId string) (models.Order, error) {
	var order models.Order
	result := utils.DB.First(&order, orderId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return order, errors.New("no outlet found for given outletId")
	}

	return order, nil
}
