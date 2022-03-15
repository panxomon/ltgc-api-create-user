package service_test

import (
	"context"
	"ltgc-api-create-user/internal/client/mocks"
	"ltgc-api-create-user/internal/service"
	data "ltgc-api-create-user/internal/service/mocks"
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_service(t *testing.T) {

	ctx := context.TODO()
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	var svc service.Service
	mockFirestore := new(mocks.Repository)

	t.Run("client ok", func(t *testing.T) {

		mockFirestore.On("CreateUser", mock.Anything, data.UserMock).Return(data.UserMock, nil)

		svc = service.MakeService(logger, mockFirestore)

		svc.CreateUser(ctx, data.UserRequest)

		assert.Equal(t, data.MockResponse.User.Name, data.UserRequest.Name)
	})
}
