package mocks

import (
	context "context"
	entity "ltgc-api-create-user/internal/entity"
)

var UserRequest *entity.Request = &entity.Request{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var UserMock *entity.User = &entity.User{
	Name:     "Pancho",
	Password: "123",
	Mail:     "francisco.mfu@lala.com",
}

var MockResponse *entity.Response = &entity.Response{
	User: *UserMock,
}

var ErrorComun = "error"

func CreateUserOK(_ context.Context, _ entity.Request) (response *entity.Response, err error) {
	return &entity.Response{
		User: *UserMock,
	}, nil
}
