package middleware

import (
	"errors"
	"json-server-kit/core"
	"net/http"

	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	DefaultCode    = http.StatusOK
	DefaultMsgEnum = utils.NewDict(map[int]string{
		http.StatusOK:                            "success",
		http.StatusBadRequest:                    "bad request",
		http.StatusUnauthorized:                  "unauthorized",
		http.StatusForbidden:                     "forbidden",
		http.StatusNotFound:                      "not found",
		http.StatusInternalServerError:           "internal server error",
		http.StatusBadGateway:                    "bad gateway",
		http.StatusServiceUnavailable:            "service unavailable",
		http.StatusGatewayTimeout:                "gateway timeout",
		http.StatusHTTPVersionNotSupported:       "http version not supported",
		http.StatusVariantAlsoNegotiates:         "variant also negotiates",
		http.StatusInsufficientStorage:           "insufficient storage",
		http.StatusLoopDetected:                  "loop detected",
		http.StatusNotExtended:                   "not extended",
		http.StatusNetworkAuthenticationRequired: "network authentication required",
		http.StatusMisdirectedRequest:            "misdirected request",
		http.StatusUnprocessableEntity:           "unprocessable entity",
		http.StatusLocked:                        "locked",
		http.StatusFailedDependency:              "failed dependency",
		http.StatusTooEarly:                      "too early",
		http.StatusUpgradeRequired:               "upgrade required",
		http.StatusPreconditionRequired:          "precondition required",
		http.StatusTooManyRequests:               "too many requests",
		http.StatusRequestHeaderFieldsTooLarge:   "request header fields too large",
		http.StatusUnavailableForLegalReasons:    "unavailable for legal reasons",
		http.StatusNotImplemented:                "not implemented",
	})
	ResponseDataKey = "responseData"
	ResponseCodeKey = "responseCode"
	ResponseMsgKey  = "responseMsg"
)

func Response() fiber.Handler {
	DefaultMsgEnum.Set(http.StatusOK, "success")
	return func(c *fiber.Ctx) error {
		response := core.NewResponse(c)
		err := c.Next()
		if err != nil {
			e := new(fiber.Error)
			if errors.As(err, &e) {
				code := e.Code
				msg := e.Message
				return response.SetCode(code).SetMsg(msg).SetData(nil).Return()
			} else {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
		}
		data := c.Locals(ResponseDataKey)
		code, ok := c.Locals(ResponseCodeKey).(int)
		if !ok {
			code = DefaultCode
		}
		msg, ok := c.Locals(ResponseMsgKey).(string)
		if !ok {
			msg = DefaultMsgEnum.Get(code)
		}
		return response.SetCode(code).SetMsg(msg).SetData(data).Return()
	}
}
