package repository

import (
	"shortener/src/model"
)

func Insert(shortener model.Shortener) (model.Shortener, error) {
	err := dbConnection.Create(&shortener).Error

	if err != nil {
		return model.Shortener{}, err
	}

	return shortener, nil
}

func FindBySlug(slug string) (model.Shortener, error) {
	var shortener model.Shortener

	err := dbConnection.Where("slug = ?", slug).First(&shortener).Error

	if err != nil {
		return model.Shortener{}, err
	}

	return shortener, nil
}
