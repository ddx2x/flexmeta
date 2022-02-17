package mongo

import (
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	metadata          = "metadata"
	version           = "version"
	metadataName      = "metadata.name"
	metadataWorkspace = "metadata.workspace"
	metadataUUID      = "metadata.uuid"
	metadataDelete    = "metadata.is_delete"
)

func connect(ctx context.Context, uri string) (*mongo.Client, error) {
	cliOpts := options.Client()
	cliOpts.SetRegistry(bson.NewRegistryBuilder().
		RegisterTypeMapEntry(
			bsontype.DateTime,
			reflect.TypeOf(time.Time{})).
		Build(),
	)
	cliOpts.ApplyURI(uri)

	mgcli, err := mongo.NewClient(cliOpts)
	if err != nil {
		return nil, err
	}

	if err := mgcli.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mgcli, nil
}

type query struct {
	DB   string         `json:"db"`
	Coll string         `json:"coll"`
	Q    map[string]any `json:"q"`
}

func parseQ(q map[string]any) *query {
	return &query{
		DB:   "base",
		Coll: "test",
		Q:    make(map[string]any),
	}
}
