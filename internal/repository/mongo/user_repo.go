package mongo

import (
	"context"
	"time"

	"github.com/monforje/dsl-edu-user/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{col: col}
}

func (r *UserRepository) GetUser(ctx context.Context, telegramID int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"telegram_id": telegramID}

	var user model.User
	err := r.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) IsExist(ctx context.Context, telegramID int64) (bool, error) {
	user, err := r.GetUser(ctx, telegramID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
