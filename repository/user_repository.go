package repository

import "go-test/entity"

type UserRepository interface {
	FindByName(name string) (*entity.User, error)
	FindById(id int) *entity.User
}
