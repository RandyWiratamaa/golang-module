package repository

import "github.com/RandyWiratamaa/golang-module/v2/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
