package endpoint

import (
	"context"
	"ltgc-api-create-user/internal/entity"

	service "ltgc-api-create-user/internal/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	CreateUser endpoint.Endpoint
}

func MakeEndpoint(s service.Service) Endpoint {
	return Endpoint{
		CreateUser: makeCreateUserEndpoint(s),
	}

}

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(entity.CreateUserRequest)

		user, err := s.Call(ctx, req.Name)

		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
