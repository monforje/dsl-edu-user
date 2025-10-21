package user

import "github.com/monforje/dsl-edu-user/internal/repository/mongo"

type Service struct {
	db *mongo.Mongo
}

func New(db *mongo.Mongo) *Service {
	return &Service{db: db}
}
