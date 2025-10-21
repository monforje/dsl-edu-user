package user

import (
	"context"
	"time"

	"github.com/monforje/dsl-edu-user/internal/repository/mongo"
)

func (s *Service) IsExist(telegramID int64) (bool, error) {
	repo := mongo.NewUserRepository(s.db.Collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.IsExist(ctx, telegramID)
}
