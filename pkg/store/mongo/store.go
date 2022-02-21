package mongo

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoCli[K, Q, R]) Create(context.Context, R) error {
	return nil
}

func (m *MongoCli[K, Q, R]) Update(ctx context.Context, old R, new R, force bool) (R, error) {
	return new, nil
}

func (m *MongoCli[K, Q, R]) List(ctx context.Context, q Q) ([]R, error) {
	var targets []R
	fOpts := options.Find()
	query := parseQ[K](q)

	cursor, err := m.cli.Database(query.DB).
		Collection(query.Coll).
		Find(ctx, query.Q, fOpts)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &targets); err != nil {
		return nil, err
	}

	for _, t := range targets {
		fmt.Println("t---->", t)
		fmt.Println("t name---->", reflect.TypeOf(t).Name())
	}

	return targets, nil
}

func (m *MongoCli[K, Q, R]) Get(ctx context.Context, q Q) (R, error) {
	var t R
	return t, nil
}

func (m *MongoCli[K, Q, R]) Delete(context.Context, Q) error {
	return nil
}
