package middleware

import (
	"errors"
	"net/http"

	utils2 "jsonix-kit/utils"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

var (
	DefaultCode    = http.StatusOK
	DefaultMsgEnum = utils.NewDict(map[int]string{
		http.StatusOK:                            "success",
		http.StatusBadRequest:                    "bad_request",
		http.StatusUnauthorized:                  "unauthorized",
		http.StatusForbidden:                     "forbidden",
		http.StatusNotFound:                      "not_found",
		http.StatusInternalServerError:           "internal_server_error",
		http.StatusBadGateway:                    "bad_gateway",
		http.StatusServiceUnavailable:            "service_unavailable",
		http.StatusGatewayTimeout:                "gateway_timeout",
		http.StatusHTTPVersionNotSupported:       "http_version_not_supported",
		http.StatusVariantAlsoNegotiates:         "variant_also_negotiates",
		http.StatusInsufficientStorage:           "insufficient_storage",
		http.StatusLoopDetected:                  "loop_detected",
		http.StatusNotExtended:                   "not_extended",
		http.StatusNetworkAuthenticationRequired: "network_authentication_required",
		http.StatusMisdirectedRequest:            "misdirected_request",
		http.StatusUnprocessableEntity:           "parameters_error",
		http.StatusLocked:                        "locked",
		http.StatusFailedDependency:              "failed_dependency",
		http.StatusTooEarly:                      "too_early",
		http.StatusUpgradeRequired:               "upgrade_required",
		http.StatusPreconditionRequired:          "precondition_required",
		http.StatusTooManyRequests:               "too_many_requests",
		http.StatusRequestHeaderFieldsTooLarge:   "request_header_fields_too_large",
		http.StatusUnavailableForLegalReasons:    "unavailable_for_legal_reasons",
		http.StatusNotImplemented:                "not_implemented",
	})
	ResponseDataKey = "responseData"
	ResponseCodeKey = "responseCode"
	ResponseMsgKey  = "responseMsg"
)

func Response() fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := utils2.NewResponse(c)
		err := c.Next()
		if err != nil {
			e := new(fiber.Error)
			if errors.As(err, &e) {
				code := e.Code
				msg := e.Message
				if localize, err := fiberi18n.Localize(c, msg); err == nil {
					msg = localize
				}
				return response.SetCode(code).SetMsg(msg).SetData(nil).Return()
			} else {
				msg := err.Error()
				if localize, err := fiberi18n.Localize(c, msg); err == nil {
					msg = localize
				}
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"message": msg,
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
		if localize, err := fiberi18n.Localize(c, msg); err == nil {
			msg = localize
		}
		return response.SetCode(code).SetMsg(msg).SetData(data).Return()
	}
}
