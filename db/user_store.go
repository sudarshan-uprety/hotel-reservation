package db

import (
	"context"
	"errors"

	"github.com/sudarshan-uprety/hotel-reservation/types"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) (map[string]string, error)
	UpdateUser(context.Context, string, *types.UpdateUserParams) (string, error)
}

type PostgresUserStore struct {
	client *gorm.DB
}

func NewPostgresUserStore(database *gorm.DB) *PostgresUserStore {
	return &PostgresUserStore{
		client: database,
	}
}

func (s *PostgresUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.client.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	if err := s.client.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *PostgresUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	created := s.client.WithContext(ctx).Create(user)
	if created.Error != nil {
		return nil, created.Error
	}
	return user, nil

}

func (s *PostgresUserStore) DeleteUser(ctx context.Context, userID string) (map[string]string, error) {
	var user types.User
	if err := s.client.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, err
	}
	deleted := s.client.WithContext(ctx).Delete(&types.User{}, userID)
	if deleted.Error != nil {
		return nil, deleted.Error
	}
	return map[string]string{"success": "User deleted successfully"}, nil
}

func (s *PostgresUserStore) UpdateUser(ctx context.Context, userID string, params *types.UpdateUserParams) (string, error) {
	var user types.User
	if err := s.client.WithContext(ctx).First(&user, userID).Error; err != nil {
		return "", err
	}
	if user.DeletedAt.Valid {
		return "", errors.New("user doesnot exists")
	}
	if params.FirstName != "" {
		user.FirstName = params.FirstName
	}
	if params.LastName != "" {
		user.LastName = params.LastName
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if err := s.client.WithContext(ctx).Save(&user).Error; err != nil {
		return "", err
	}
	return "user updated successfully", nil
}
