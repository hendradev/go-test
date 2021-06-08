package repository

import (
	"go-test/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindByName(name string) (*entity.User, error) {
	argument := repository.Mock.Called(name)
	if argument.Get(0) == nil {
		return nil, argument.Error(1)
	}
	user := argument.Get(0).(entity.User)
	return &user, nil
}

func (repository *UserRepositoryMock) FindById(id int) *entity.User {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	}
	user := argument.Get(0).(entity.User)
	return &user
}
