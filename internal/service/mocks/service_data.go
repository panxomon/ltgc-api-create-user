package mocks

import (
	context "context"
	entity "ltgc-api-create-user/internal/entity"
)

var UserRequest *entity.CreateUserRequest = &entity.CreateUserRequest{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var UserMock *entity.User = &entity.User{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var MockResponse *entity.CreateUserResponse = &entity.CreateUserResponse{
	User: *UserMock,
}

var ErrorComun = "error"

func CreateUserOK(_ context.Context, _ entity.CreateUserRequest) (response *entity.CreateUserResponse, err error) {
	return &entity.CreateUserResponse{
		User: *UserMock,
	}, nil
}
