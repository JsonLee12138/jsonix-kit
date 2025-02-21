package core

import (
	"context"
	"fmt"
	"time"

	"json-server-kit/configs"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/go-redis/redis/v8"
)

const (
	CaptchaKey = "captcha:%s"
)

type CaptchaStore struct {
	client *redis.Client
	ctx    context.Context
	cnf    *configs.CaptchaConfig
}

func NewCaptchaStore(client *redis.Client, cnf *configs.Config) *CaptchaStore {
	return &CaptchaStore{
		client: client,
		ctx:    context.Background(),
		cnf:    &cnf.CaptchaConfig,
	}
}

func (s *CaptchaStore) genKey(id string) string {
	return fmt.Sprintf(CaptchaKey, id)
}

func (s *CaptchaStore) Set(id, value string) error {
	key := s.genKey(id)
	expireTime := time.Duration(utils.DefaultIfEmpty[int](s.cnf.MaxAge, 5)) * time.Minute
	return s.client.Set(s.ctx, key, value, expireTime).Err()
}

func (s *CaptchaStore) Get(id string, clear bool) (value string) {
	key := s.genKey(id)
	val, err := s.client.Get(s.ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(s.ctx, key)
	}
	return val
}

func (s *CaptchaStore) Verify(id, answer string, clear bool) bool {
	key := s.genKey(id)
	val, err := s.client.Get(s.ctx, key).Result()
	if err != nil {
		return false
	}
	if val != answer {
		return false
	}
	if clear {
		s.client.Del(s.ctx, key)
	}
	return true
}
