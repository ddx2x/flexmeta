package mongo

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/dict"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoCli[K, Q, R]) Create(context.Context, R) error {
	return nil
}

func (m *MongoCli[K, Q, R]) Update(ctx context.Context, old R, new R, q Q) (R, error) {
	query := parseQ(q)
	singleResult := m.cli.Database(query.DB).
		Collection(query.Coll).
		FindOne(ctx, query.Q)

	if singleResult.Err() == mongo.ErrNoDocuments {
		_, err := m.cli.Database(query.DB).
			Collection(query.Coll).
			InsertOne(ctx, new)
		if err != nil {
			return new, err
		}
		return new, nil
	}

	if err := singleResult.Decode(old); err != nil {
		return new, err
	}

	oldObject, newObject := &core.Object[R]{Item: old}, &core.Object[R]{Item: new}

	oldMap, err := oldObject.ToMap()
	if err != nil {
		return new, err
	}

	newMap, err := newObject.ToMap()
	if err != nil {
		return new, err
	}

	for _, path := range query.Paths {
		dict.CompareMergeObject(oldMap, newMap, path)
	}

	if err := newObject.From(oldMap); err != nil {
		return old, err
	}

	upsert := true
	_, err = m.cli.Database(query.DB).
		Collection(query.Coll).
		ReplaceOne(ctx,
			query.Q,
			new,
			options.MergeReplaceOptions(
				&options.ReplaceOptions{Upsert: &upsert},
			),
		)
	if err != nil {
		return new, err
	}

	return new, nil
}

func (m *MongoCli[K, Q, R]) List(ctx context.Context, q Q) ([]R, error) {
	var targets []R
	fOpts := options.Find()
	query := parseQ(q)

	cursor, err := m.cli.Database(query.DB).
		Collection(query.Coll).
		Find(ctx, query.Q, fOpts)
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
	query := parseQ(q)
	singleResult := m.cli.Database(query.DB).
		Collection(query.Coll).
		FindOne(ctx, query)
	if err := singleResult.Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

func (m *MongoCli[K, Q, R]) Delete(context.Context, Q) error {
	return nil
}
