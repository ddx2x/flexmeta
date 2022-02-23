package mongo

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NotFound = fmt.Errorf("notFound")

const (
	metadata          = "metadata"
	version           = "version"
	metadataName      = "metadata.name"
	metadataWorkspace = "metadata.workspace"
	metadataUUID      = "metadata.uuid"
	metadataDelete    = "metadata.deleted"
)

func getCtx(client *mongo.Client) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, cancel, err
	}
	return ctx, cancel, nil
}

func connect(ctx context.Context, uri string) (*mongo.Client, error) {
	cliOpt := options.Client()
	cliOpt.SetRegistry(
		bson.NewRegistryBuilder().
			RegisterTypeMapEntry(
				bsontype.DateTime,
				reflect.TypeOf(time.Time{})).
			Build(),
	)
	cliOpt.ApplyURI(uri)
	mcli, err := mongo.NewClient(cliOpt)
	if err != nil {
		return nil, err
	}
	ctx, cancel, err := getCtx(mcli)
	defer cancel()
	if err != nil {
		return nil, err
	}
	if err := mcli.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return mcli, nil
}

type query struct {
	DB    string   `json:"db"`
	Coll  string   `json:"coll"`
	Paths []string `json:"paths"`
	Q     bson.D   `json:"q"`
}

func parseQ[K comparable, Q ~map[K]any](q Q) *query {
	return &query{
		DB:   "base",
		Coll: "account",
		Q:    bson.D{},
	}
}

func versionMatchFilter(opData map[string]interface{}, resourceVersion string) bool {
	if resourceVersion == "" {
		return false
	}
	md, exist := opData[metadata]
	if !exist {
		return false
	}
	mdMap := md.(map[string]interface{})
	version, exist := mdMap[version]
	if !exist {
		return false
	}
	if version.(string) <= resourceVersion {
		return false
	}
	return true
}
