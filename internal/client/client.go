package client

import (
	"context"
	"log"
	"ltgc-api-create-user/internal/entity"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Repository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type firestoreRepository struct {
	collection string
	client     *firestore.Client
	timeout    time.Duration
}

func newClient(file string, projectID string, timeout int) (*firestore.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	sa := option.WithCredentialsFile(file)
	client, err := firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewFirestorerepository(fileName string, projectID string, collectionName string) Repository {
	client, err := newClient(fileName, projectID, 10)
	if err != nil {
		log.Println("error creando un cliente firestore %w", err)
	}
	r := &firestoreRepository{
		collection: collectionName,
		client:     client,
		timeout:    time.Duration(10) * time.Second,
	}

	return r
}

func (r *firestoreRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {

	_, err := r.client.Collection(r.collection).Doc(user.Name).Set(ctx, user)
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}
