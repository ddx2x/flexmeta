package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoCli[K, Q, R]) Create(ctx context.Context, q Q, r R) (R, error) {
	return r, nil
}

func (m *MongoCli[K, Q, R]) Update(ctx context.Context, q Q, old R, new R, force bool) (R, error) {
	return new, nil
}

func (m *MongoCli[K, Q, R]) List(ctx context.Context, q Q) ([]R, error) {
	var targets []R
	fOpts := options.Find()
	query := parseQ(q)

	cursor, err := m.cli.Database(query.DB).Collection(query.Coll).Find(ctx, query.Q, fOpts)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &targets); err != nil {
		return nil, err
	}

	return targets, nil
}

func (m *MongoCli[K, Q, R]) Get(ctx context.Context, q Q) (R, error) {
	var t R
	return t, nil
}
