package repository

import "github.com/RandyWiratamaa/golang-module/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
