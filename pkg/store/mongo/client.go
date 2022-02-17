package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCli[K string, Q map[string]any, R any] struct {
	cli *mongo.Client
}

func (m *MongoCli[K, Q, R]) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.cli.Disconnect(ctx)
}

func NewMongoCli[K string, Q map[string]any, R any](ctx context.Context, uri string) (*MongoCli[K, Q, R], error) {
	cli, err := connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	return &MongoCli[K, Q, R]{cli: cli}, nil
}
