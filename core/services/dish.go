package services

import (
	"errors"
	"skipthequeue/core/models"

	"skipthequeue/utils"

	"gorm.io/gorm"
)

func FindDishByIdService(dishId string) (models.Dish, error) {
	var dish models.Dish
	result := utils.DB.First(&dish, dishId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return dish, errors.New("no outlet found for given outletId")
	}

	return dish, nil
}
