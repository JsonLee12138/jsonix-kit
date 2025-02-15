package repository

import (
	"context"
	"errors"
	"fmt"
	"json-server-kit/apps/common/entity"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

const (
	PasswordErrorMaxAge = time.Minute * 5
)

func NewUserRepository(db *gorm.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{
		db:    db,
		redis: redis,
	}
}

func (r *UserRepository) GetUserInfoByUsername(username string) (*entity.User, error) {
	var userInfo entity.User
	if err := r.db.Where("username = ?", username).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (r *UserRepository) GetUserInfoByPhone(phone string) (*entity.User, error) {
	var userInfo entity.User
	if err := r.db.Where("phone = ?", phone).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (r *UserRepository) GetUserInfoByEmail(email string) (*entity.User, error) {
	var userInfo entity.User
	if err := r.db.Where("email = ?", email).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (r *UserRepository) GetUserInfoByAccount(account string) (*entity.User, error) {
	var userInfo entity.User
	if err := r.db.Where("username = ? OR phone = ? OR email = ?", account, account, account).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func UserPasswordErrorKey(id string) string {
	return fmt.Sprintf("password_error:%s", id)
}

func (r *UserRepository) SetPasswordError(id string) (int64, error) {
	key := UserPasswordErrorKey(id)
	count, err := r.redis.Incr(context.Background(), key).Result()
	if count == 1 {
		r.redis.Expire(context.Background(), key, PasswordErrorMaxAge)
	}
	return count, err
}

func (r *UserRepository) ClearPasswordError(id string) error {
	return r.redis.Del(context.Background(), UserPasswordErrorKey(id)).Err()
}

func (r *UserRepository) GetPasswordError(id string) (int64, error) {
	return r.redis.Get(context.Background(), UserPasswordErrorKey(id)).Int64()
}

func (r *UserRepository) DisableUserForever(id string) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("enable", false).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
