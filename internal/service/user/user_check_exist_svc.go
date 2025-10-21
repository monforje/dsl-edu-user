package user

import (
	"context"

	"github.com/monforje/dsl-edu-user/internal/repository/mongo"
)

func (s *Service) IsExist(telegramID int64) (bool, error) {
	repo := mongo.NewUserRepository(s.db.Collection)
	return repo.IsExist(context.Background(), telegramID)
}
