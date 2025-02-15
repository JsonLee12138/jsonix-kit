package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

//go:generate easyjson -all ./logger.go

type LogVO struct {
	Time      time.Time
	Path      string
	Query     string
	Body      string
	Spend     time.Duration
	IP        string
	UserAgent string
	Error     string
	Source    string
	Headers   map[string][]string
	Method    string
	Status    int
	Code      int
}

func Logger(print func(vo LogVO)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()
		path := c.Path()
		query := c.Request().URI().QueryString()
		body := c.Request().Body()
		c.Request().SetBody(body)
		spend := time.Since(startTime)
		headers := c.GetReqHeaders()
		userAgent := c.Get("User-Agent")
		method := c.Method()
		c.Next()
		data := LogVO{
			Time:      startTime,
			Path:      path,
			Query:     string(query),
			Body:      string(body),
			Spend:     spend,
			IP:        c.IP(),
			UserAgent: userAgent,
			Headers:   headers,
			Method:    method,
			Status:    c.Response().StatusCode(),
		}
		code := c.Get("code")
		codeInt, err := strconv.Atoi(code)
		if err != nil {
			codeInt = http.StatusOK
		}
		data.Code = codeInt
		print(data)
		return nil
	}
}
