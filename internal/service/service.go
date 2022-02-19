package service

import (
	"context"
	"ltgc-api-create-user/internal/entity"

	"github.com/go-kit/kit/log"
)

type Service interface {
	Call(ctx context.Context, name string) (*entity.User, error)
}

type service struct {
	logger     log.Logger
	repository entity.Repository
}

func MakeService(log log.Logger, r entity.Repository) Service {
	return &service{
		logger:     log,
		repository: r,
	}
}

func (s service) Call(ctx context.Context, name string) (*entity.User, error) {

	user, err := s.repository.CreateUser(ctx, name)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
