package endpoint

import (
	"context"

	"ltgc-api-create-user/internal/entity"
	"ltgc-api-create-user/internal/service"

	"github.com/go-kit/kit/endpoint"
)

// type Endpoint struct {
// 	CreateUser epkit.Endpoint
// }

// func MakeEndpoints(s service.Service, logger log.Logger) epkit.Endpoint {
// 	return Endpoint{
// 		CreateUser: makeCreateUserEndpoint(s, logger),
// 	}

// }

// func makeCreateUserEndpoint(svc service.Service, logger log.Logger) epkit.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {

// 		req, ok := request.(entity.CreateUserRequest)
// 		if !ok {
// 			level.Error(logger).Log("message", "invalid request")
// 			return nil, errors.New("invalid request")
// 		}

// 		user, err := svc.CreateUser(ctx, req)

// 		if err != nil {
// 			return nil, err
// 		}
// 		return user, nil
// 	}
// }

func MakeServiceEndpoint(ctx context.Context, svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*entity.CreateUserRequest)

		return svc.CreateUser(ctx, req)
	}
}
