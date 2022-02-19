package entity

import "context"

type User struct {
	Id       int64
	Name     string
	Mail     string
	Password string
}

type Repository interface {
	CreateUser(ctx context.Context, name string) (User, error)
}
