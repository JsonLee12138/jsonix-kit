package utils

import (
	"errors"
	"json-server-kit/configs"
	"time"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/xid"
	"golang.org/x/sync/singleflight"
)

type UserClaims struct {
	ID       xid.ID
	Username string
	Nickname string
	Phone    string
	Email    string
}

type BaseClaims struct {
	UserClaims UserClaims
	BufferTime *jwt.NumericDate
	DeviceId   string
	jwt.RegisteredClaims
}

type JWT struct {
	Config             configs.JWTConfig
	ConcurrencyControl *singleflight.Group
}

func NewJWT(config configs.JWTConfig) *JWT {
	if utils.IsEmpty(config.SigningMethod) {
		config.SigningMethod = "HS256"
	}
	if !utils.IsEmpty(config.SecretKey) {
		config.SecretKey = utils.Sha256String(config.SecretKey)
	}
	return &JWT{
		Config:             config,
		ConcurrencyControl: &singleflight.Group{},
	}
}

func (j *JWT) GetExpireTime() time.Duration {
	expireTime, _ := utils.ParseDuration(j.Config.ExpireTime)
	return expireTime
}
func (j *JWT) GetBufferTime() time.Duration {
	bufferTime, _ := utils.ParseDuration(j.Config.BufferTime)
	return bufferTime
}
func (j *JWT) CreateClaims(claims BaseClaims) BaseClaims {
	expireTime, _ := utils.ParseDuration(j.Config.ExpireTime)
	bufferTime, _ := utils.ParseDuration(j.Config.BufferTime)

	return BaseClaims{
		UserClaims: claims.UserClaims,
		BufferTime: jwt.NewNumericDate(time.Now().Add(bufferTime)),
		DeviceId:   claims.DeviceId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1 * time.Second)),
		},
	}
}

func (j *JWT) CreateToken(claims BaseClaims) (string, error) {
	signinMethod := jwt.GetSigningMethod(j.Config.SigningMethod)
	token := jwt.NewWithClaims(signinMethod, claims)
	return token.SignedString([]byte(j.Config.SecretKey))
}

func (j *JWT) ParseToken(value string) (*BaseClaims, error) {
	token, err := jwt.ParseWithClaims(value, &BaseClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.Config.SecretKey), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, errors.New("token_malformed")
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, errors.New("token_expired")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, errors.New("token_active_not_yet")
		default:
			return nil, errors.New("token_invalid")
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*BaseClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token_invalid")
	}
	return nil, errors.New("token_invalid")
}

func (j *JWT) VerifyToken(value string) (*BaseClaims, error) {
	return j.ParseToken(value)
}

func (j *JWT) ExchangeToken(old string, claims BaseClaims) (string, error) {
	v, err, _ := j.ConcurrencyControl.Do("JWT:"+old, func() (any, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}
