package mongo

import (
	"context"
	"fmt"
	"reflect"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/dict"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var nilfilter = map[string]any{}

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

func fieldMatchFilter(opData map[string]interface{}, key string, value interface{}) bool {
	return reflect.DeepEqual(dict.Get(opData, key), value)
}

func (m *MongoCli[K, Q, R]) Watch(ctx context.Context, q Q) (<-chan core.Event, <-chan error) {
	errC := make(chan error)
	query := parseQ(q)

	ns := fmt.Sprintf("%s.%s", query.DB, query.Coll)
	directReadFilter := func(op *Op) bool {
		pass := true
		for _, filter := range query.Q {
			if pass = fieldMatchFilter(op.Data, filter.Key, filter.Value); !pass {
				pass = false
				break
			}
		}
		return true
	}
	oplogTailCtx := Start(m.cli, &Options{
		DirectReadNs:     []string{ns},
		ChangeStreamNs:   []string{ns},
		MaxAwaitTime:     10,
		DirectReadFilter: directReadFilter,
	})

	eventC := make(chan core.Event, 0)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(eventC)
				oplogTailCtx.Stop()
				return
			case <-oplogTailCtx.ErrC:
				close(eventC)
				return
			case op, ok := <-oplogTailCtx.OpC:
				if !ok {
					return
				}
				var evtop core.EventType
				switch {
				case op.IsInsert():
					evtop = core.ADDED
					if isDelete := dict.Get(op.Data, DELETED); isDelete != nil {
						if v, ok := isDelete.(bool); ok && v {
							continue
						}
					}
				case op.IsUpdate():
					evtop = core.MODIFIED
					if isDelete := dict.Get(op.Data, DELETED); isDelete != nil {
						if v, ok := isDelete.(bool); ok && v {
							evtop = core.DELETED
						}
					}
				case op.IsDelete():
					evtop = core.DELETED
				}
				var r R
				if err := r.Unmarshal(op.Data); err != nil {
					errC <- err
					return
				}
				evt := core.Event{
					Type:   evtop,
					Object: r,
				}
				eventC <- evt
			}
		}
	}()

	return eventC, errC
}

func (m *MongoCli[K, Q, R]) checkExistAndCreate(ctx context.Context, db, collection string, enableIndex bool, uniqeKeys ...string) error {
	names, err := m.cli.Database(db).
		ListCollectionNames(ctx, nilfilter)
	if err != nil {
		return err
	}
	exist := false

	for _, name := range names {
		if name == collection {
			exist = true
		}
	}

	if !exist {
		if err := m.cli.Database(db).CreateCollection(ctx, collection); err != nil {
			return err
		}
	}

	if enableIndex {
		keys := bson.D{}
		for _, k := range uniqeKeys {
			keys = append(keys, bson.E{Key: k, Value: 1})
		}
		indexModel := mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(true),
		}
		if _, err = m.cli.Database(db).
			Collection(collection).
			Indexes().
			CreateOne(ctx, indexModel); err != nil {
			return err
		}
	}

	return nil
}
