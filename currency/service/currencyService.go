package service

import (
	"currency/db"
	"currency/models"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

func FindAll() ([]*models.CurrencyResponse, error) {
	var all []models.Currency
	result := db.DB.Find(&all)
	if result.Error != nil {
		return nil, result.Error
	}

	var response []*models.CurrencyResponse
	err := mapstructure.Decode(all, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func FindById(id int) (*models.CurrencyResponse, error) {
	var currency models.Currency
	err := db.DB.Where("id = ?", id).First(&currency).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var response models.CurrencyResponse
	err = mapstructure.Decode(currency, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Create(request *models.CurrencyRequest) (*models.CurrencyResponse, error) {
	var currency models.Currency
	err := mapstructure.Decode(request, &currency)
	if err != nil {
		return nil, err
	}
	err = db.DB.Create(&currency).Error
	if err != nil {
		return nil, err
	}

	var response models.CurrencyResponse
	err = mapstructure.Decode(currency, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Update(id int, request *models.CurrencyRequest) (*models.CurrencyResponse, error) {
	var currency models.Currency

	result := db.DB.First(&currency, id)
	if result.Error != nil {
		return nil, result.Error
	}

	err := mapstructure.Decode(request, &currency)
	if err != nil {
		return nil, err
	}

	result = db.DB.Save(&currency)
	if result.Error != nil {
		return nil, result.Error
	}

	var response models.CurrencyResponse
	err = mapstructure.Decode(currency, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func Delete(id int) error {
	var currency models.Currency

	result := db.DB.First(&currency, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.DB.Delete(&currency)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
