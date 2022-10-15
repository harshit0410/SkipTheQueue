package services

import (
	"errors"
	"skipthequeue/pkg/models"

	"skipthequeue/utils"

	"gorm.io/gorm"
)

func FindOutletByIdService(outletId string) (models.Outlet, error) {
	var outlet models.Outlet
	result := utils.DB.First(&outlet, outletId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return outlet, errors.New("no outlet found for given outletId")
	}

	return outlet, nil
}
