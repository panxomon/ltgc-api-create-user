package client

import "context"

type Firestore interface {
	Execute(ctx context.Context)
}

type client struct {
	firestore Firestore
}

