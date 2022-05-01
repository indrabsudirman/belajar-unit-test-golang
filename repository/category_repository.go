package repository

import "belajar-unit-test-golang/entity"

type CategoryRepository interface {
	//Kontrak ke category entity
	FindById(id string) *entity.Category
}
