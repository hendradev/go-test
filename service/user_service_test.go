package service

import (
	"fmt"
	"go-test/entity"
	"go-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestGetSuccess(t *testing.T) {
	userMock := entity.User{Name: "Andi"}
	userRepository.Mock.On("FindByName", "Andi").Return(userMock, nil)
	user := userService.Get("Andi")
	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "Andi")
}

func TestGetByIdSuccess(t *testing.T) {
	userMock := entity.User{Name: "Andi"}
	userRepository.Mock.On("FindById", mock.AnythingOfType("int")).Return(userMock)
	user := userService.GetById(1)
	assert.NotNil(t, user)
	assert.Equal(t, user.Name, "Andi")
}

func TestGetByIdArgumentNil(t *testing.T) {
	userRepositoryB := &repository.UserRepositoryMock{Mock: mock.Mock{}}
	userServiceB := UserService{Repository: userRepositoryB}
	userRepositoryB.Mock.On("FindById", 0).Return(nil)
	user := userServiceB.GetById(0)
	fmt.Println(user)
	assert.Nil(t, user)
}
