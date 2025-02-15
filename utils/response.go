package utils

import (
	"math"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	timeFormat = time.RFC3339
)

type IResponseInterface interface {
	NewListResponse(data *ListVO[any]) *Response
	NewSuccess(data any) *Response
	NewError(msg string) *Response
	SetCode(code int) *Response
	SetMsg(msg string) *Response
	NewFailWithParse() *Response
	NewBadGateway() *Response
	NewForbidden() *Response
	NewRefreshTokenInvalid() *Response
	NewUnauthorized() *Response
	SetHeaderOfStream() *Response
	FailWithParse() error
	Forbidden() error
	Unauthorized() error
	RefreshTokenInvalid() error
	Return()
}

type BaseResponseVO struct {
	Code int
	Msg  string
	Data any
	Time time.Time
}

type Response struct {
	VO      *BaseResponseVO
	Context *fiber.Ctx
}

func NewResponse(ctx *fiber.Ctx) *Response {
	return &Response{
		Context: ctx,
		VO:      new(BaseResponseVO),
	}
}

type ListVO[T any] struct {
	List      []T   `json:"list"`
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPages"`
	HasMore   bool  `json:"hasMore"`
}

func NewList[T any](data *ListVO[T]) *ListVO[T] {
	data.TotalPage = int(math.Ceil(float64(data.Total) / float64(data.PageSize)))
	data.HasMore = data.TotalPage > data.Page
	return data
}

func (r *Response) NewListResponse(data *ListVO[any]) *Response {
	data = NewList(data)
	r.VO = &BaseResponseVO{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
	return r
}

func (r *Response) NewSuccess(data any) *Response {
	r.VO = &BaseResponseVO{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
	return r
}

func (r *Response) NewError(msg string) *Response {
	if msg == "" {
		msg = "error"
	}
	r.VO = &BaseResponseVO{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	}
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.VO.Code = code
	return r
}

func (r *Response) SetMsg(msg string) *Response {
	r.VO.Msg = msg
	return r
}

func (r *Response) SetData(data any) *Response {
	r.VO.Data = data
	return r
}

func (r *Response) NewBadParameters() *Response {
	return r.NewError("response.fail_with_parse").SetCode(http.StatusUnprocessableEntity)
}
func (r *Response) NewBadGateway() *Response {
	return r.NewError("response.bad_gateway").SetCode(http.StatusBadGateway)
}

func (r *Response) NewForbidden() *Response {
	return r.NewError("response.forbidden").SetCode(http.StatusForbidden)
}

func (r *Response) NewRefreshTokenInvalid() *Response {
	r.Context.Set("Refresh-Token-Invalid", "true")
	return r.NewError("response.token_invalid").SetCode(http.StatusForbidden)
}

func (r *Response) NewUnauthorized() *Response {
	return r.NewError("response.unauthorized").SetCode(http.StatusUnauthorized)
}

func (r *Response) SetHeaderOfStream() *Response {
	r.Context.Set("Content-Type", "text/event-stream; charset=utf-8")
	r.Context.Set("Cache-Control", "no-cache")
	r.Context.Set("Connection", "keep-alive")
	r.Context.Set("Transfer-Encoding", "chunked")
	return r
}
func (r *Response) BadParameters() error {
	return r.NewBadGateway().Return()
}
func (r *Response) Forbidden() error {
	return r.NewForbidden().Return()
}
func (r *Response) Unauthorized() error {
	return r.NewUnauthorized().Return()
}
func (r *Response) BadGateway() error {
	return r.NewBadGateway().Return()
}
func (r *Response) RefreshTokenInvalid() error {
	return r.NewRefreshTokenInvalid().Return()
}
func (r *Response) Return() error {
	r.Context.Locals("code", r.VO.Code)
	return r.Context.Status(http.StatusOK).JSON(fiber.Map{
		"code": r.VO.Code,
		"msg":  r.VO.Msg,
		"data": r.VO.Data,
		"time": time.Now().Format(timeFormat),
	})
}

func NewBadParameters(errs ...error) *fiber.Error {
	var err string
	if len(errs) == 0 {
		err = "parameters_error"
	} else {
		err = errs[0].Error()
	}
	return fiber.NewError(http.StatusUnprocessableEntity, err)
}
