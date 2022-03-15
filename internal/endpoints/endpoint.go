package endpoint

import (
	"context"

	"ltgc-api-create-user/internal/entity"
	"ltgc-api-create-user/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeServiceEndpoint(ctx context.Context, svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*entity.CreateUserRequest)

		return svc.CreateUser(ctx, req)
	}
}
