package service

import (
	"errors"

	"github.com/RandyWiratamaa/golang-module/v2/entity"
	"github.com/RandyWiratamaa/golang-module/v2/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
