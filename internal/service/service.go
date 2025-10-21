package service

import (
	"github.com/monforje/dsl-edu-user/internal/repository/mongo"
	mongosvc "github.com/monforje/dsl-edu-user/internal/service/user"
	"github.com/monforje/dsl-edu-user/internal/transport/http"
)

type Service struct {
	Mongo *mongosvc.Service
}

func New(db *mongo.Mongo /*остальные сервисы*/) *Service {
	return &Service{
		Mongo: mongosvc.New(db),
	}
}

func (s *Service) MongoService() http.MongoService {
	return s.Mongo
}
