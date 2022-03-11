package mocks

import (
	entity "ltgc-api-create-user/internal/entity"
)

var UserRequest *entity.CreateUserRequest = &entity.CreateUserRequest{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var MockResponse *entity.User = &entity.User{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var ErrorComun = "error"
