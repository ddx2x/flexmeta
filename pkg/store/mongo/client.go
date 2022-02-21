package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var cli *mongo.Client

type MongoCli[K comparable, Q ~map[K]any, R any] struct {
	cli *mongo.Client
}

func (m *MongoCli[K, Q, R]) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.cli.Disconnect(ctx)
}

func NewMongoCli[K comparable, Q ~map[K]any, R any](ctx context.Context, uri string) (*MongoCli[K, Q, R], error) {
	var err error
	if cli == nil {
		cli, err = connect(ctx, uri)
		if err != nil {
			return nil, fmt.Errorf("connect to mongo: %w", err)
		}
	}
	return &MongoCli[K, Q, R]{
		cli: cli,
	}, nil
}
