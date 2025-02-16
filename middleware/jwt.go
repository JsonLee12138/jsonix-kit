package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"json-server-kit/apps/auth/model/enum"
	"json-server-kit/configs"
	"json-server-kit/utils"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
)

type JWTAuthConfig struct {
	RefreshCnf configs.JWTConfig
	AccessCnf  configs.JWTConfig
	Redis      *redis.Client
	Uaparser   *uaparser.Parser
}

func genRefreshKey(deviceId string, userId string, client *uaparser.Client) string {
	clientType := enum.ClientTypeWeb
	if enum.IsApp(client.UserAgent.Family) {
		clientType = enum.ClientTypeApp
	}
	return fmt.Sprintf("refresh_token:%s:%s:%s", deviceId, clientType, userId)
}

func genAccessTokenKey(deviceId string, userId string, client *uaparser.Client) string {
	clientType := enum.ClientTypeWeb
	if enum.IsApp(client.UserAgent.Family) {
		clientType = enum.ClientTypeApp
	}
	return fmt.Sprintf("access_token:%s:%s:%s", deviceId, clientType, userId)
}

func JWTAuth(cnf JWTAuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("Authorization")
		refreshToken := c.Cookies("refresh_token")
		if accessToken == "" || refreshToken == "" {
			c.Locals("code", http.StatusForbidden)
			return errors.New("unauthorized")
		}
		var deviceId string
		var client *uaparser.Client
		accessClaims, err := utils.NewJWT(cnf.AccessCnf).VerifyToken(accessToken)
		if err != nil || accessClaims.ExpiresAt.Unix() < time.Now().Unix() {
			refreshClaims, err := utils.NewJWT(cnf.RefreshCnf).VerifyToken(refreshToken)
			if err != nil {
				return fiber.NewError(http.StatusForbidden, "unauthorized")
			}
			if refreshClaims.ExpiresAt.Unix() < time.Now().Unix() {
				return fiber.NewError(http.StatusForbidden, "unauthorized")
			}
			if cnf.Redis != nil {
				deviceId = c.Get("X-Device-Id")
				client = cnf.Uaparser.Parse(c.Get("User-Agent"))
				refreshKey := genRefreshKey(deviceId, refreshClaims.UserClaims.ID.String(), client)
				refreshRedisToken, err := cnf.Redis.Get(context.Background(), refreshKey).Result()
				if err != nil {
					return fiber.NewError(http.StatusForbidden, "unauthorized")
				}
				if refreshRedisToken != refreshToken {
					return fiber.NewError(http.StatusForbidden, "unauthorized")
				}
			}
			if refreshClaims.ExpiresAt.Unix() > time.Now().Unix() && refreshClaims.BufferTime.Unix() < time.Now().Unix() {
				c.Set("X-Refresh-Soon", "1")
			}
		}
		if cnf.Redis != nil {
			if deviceId == "" {
				deviceId = c.Get("X-Device-Id")
			}
			if client == nil {
				client = cnf.Uaparser.Parse(c.Get("User-Agent"))
			}
			accessKey := genAccessTokenKey(deviceId, accessClaims.UserClaims.ID.String(), client)
			accessRedisToken, err := cnf.Redis.Get(context.Background(), accessKey).Result()
			if err != nil {
				c.Locals("code", http.StatusUnauthorized)
				return errors.New("unauthorized")
			}
			if accessRedisToken != accessToken {
				c.Locals("code", http.StatusUnauthorized)
				return errors.New("unauthorized")
			}
		}
		if accessClaims.ExpiresAt.Unix() > time.Now().Unix() && accessClaims.BufferTime.Unix() < time.Now().Unix() {
			c.Set("X-Access-Soon", "1")
		}
		c.Locals("baseUserInfo", accessClaims.UserClaims)
		c.Next()
		return nil
	}
}
