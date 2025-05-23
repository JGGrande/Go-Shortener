package service

import (
	"shortener/src/model"
	"shortener/src/repository"
)

func Create(shortener model.Shortener) (model.Shortener, error) {
	// if shortener.Slug == "" {
	// 	shortener.Slug = generateNanoID()
	// }

	createdShortener, err := repository.Insert(shortener)

	if err != nil {
		return model.Shortener{}, err
	}

	return createdShortener, nil
}

func FindBySlug(slug string) (model.Shortener, error) {
	shortener, err := repository.FindBySlug(slug)

	if err != nil {
		return model.Shortener{}, err
	}

	return shortener, nil
}
