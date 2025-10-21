package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/monforje/dsl-edu-user/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

func New(cfg *config.ConfigDatabase) (*Mongo, error) {
	if cfg == nil {
		return nil, fmt.Errorf("mongo конфиг пустой")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB не отвечает: %w", err)
	}

	db := client.Database(cfg.DBname)
	col := db.Collection(cfg.COLname)

	log.Printf("MongoDB подключен: %s/%s (%s)", cfg.MongoURI, cfg.DBname, cfg.COLname)

	return &Mongo{
		Client:     client,
		DB:         db,
		Collection: col,
	}, nil
}
