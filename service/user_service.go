package service

import (
	"go-test/entity"
	"go-test/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(name string) *entity.User {
	user, err := service.Repository.FindByName(name)
	if err != nil {
		return nil
	}
	return user
}

func (service UserService) GetById(id int) *entity.User {
	user := service.Repository.FindById(id)
	return user
}
