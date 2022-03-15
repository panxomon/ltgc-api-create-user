package service

import (
	"context"
	"ltgc-api-create-user/internal/client"
	"ltgc-api-create-user/internal/entity"

	"github.com/go-kit/kit/log"
)

type Service interface {
	CreateUser(ctx context.Context, user *entity.CreateUserRequest) (*entity.User, error)
}

type service struct {
	logger     log.Logger
	repository client.Repository
}

func MakeService(logger log.Logger, repository client.Repository) Service {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) CreateUser(ctx context.Context, req *entity.CreateUserRequest) (*entity.User, error) {

	user := &entity.User{
		Name:     req.Name,
		Mail:     req.Mail,
		Password: req.Password,
	}

	userEntity, err := s.repository.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return userEntity, nil
}
