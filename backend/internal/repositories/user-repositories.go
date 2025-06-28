package repositories

import (
	"auth-jwt/internal/database"
	"auth-jwt/internal/models"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	collection *mongo.Collection // users
}

func NewUserRepository(db *database.Service) *UserRepository {
	return &UserRepository{
		collection: db.Database.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid object ID format: %v", err)
	}

	err = r.collection.FindOne(ctx, bson.D{{"_id", objectID}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	if user.ID.IsZero() {
		return errors.New("user ID is required")
	}

	user.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.D{{"_id", user.ID}}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}

	return nil
}
