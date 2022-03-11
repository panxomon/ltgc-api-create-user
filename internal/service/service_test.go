package service_test

import (
	"context"
	"ltgc-api-create-user/internal/client/mocks"
	"ltgc-api-create-user/internal/service"
	data "ltgc-api-create-user/internal/service/mocks"
	"os"
	"testing"

	"github.com/go-kit/log"
)

func Test_service(t *testing.T) {

	ctx := context.TODO()
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	var svc service.Service

	mock := new(mocks.Repository)
	mock.On("CreateUser", ctx, &data.UserRequest).Return(*data.MockResponse, nil)

	svc = service.MakeService(logger, mock)

	svc.CreateUser(ctx, data.UserRequest)

}
