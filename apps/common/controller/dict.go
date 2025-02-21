package controller

import (
	"errors"
	"json-server-kit/apps/common/model/dto"
	"json-server-kit/apps/common/service"
	"json-server-kit/utils"

	"github.com/gofiber/fiber/v2"
)

type DictController struct {
	DictService *service.DictService
}

func NewDictController(DictService *service.DictService) *DictController {
	return &DictController{
		DictService,
	}
}

// @Summary 获取字典类型
// @Description 获取字典类型
// @Tags common
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Param code query string false "code"
// @Param name query string false "name"
// @Success 200 {object} utils.BaseResponseVO{data=[]entity.DictType}
// @Router /common/dict/types [get]
func (c *DictController) GetDictTypes(ctx *fiber.Ctx) error {
	var params dto.DictTypeQuery
	if err := ctx.QueryParser(&params); err != nil {
		return utils.NewBadParameters()
	}
	dictTypes, err := c.DictService.GetDictTypes(params)
	if err != nil {
		return errors.New(err.Error())
	}
	ctx.Locals("data", dictTypes)
	return nil
}

// @Summary 获取字典值列表
// @Description 获取字典值列表
// @Tags common
// @Accept json
// @Produce json
// @Param code query string false "code"
// @Success 200 {object} utils.BaseResponseVO{data=[]entity.DictItem}
// @Router /common/dict/values [get]
func (c *DictController) GetDictValues(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	if code == "" {
		return utils.NewBadParameters()
	}
	dictItems, err := c.DictService.GetDictItems(code)
	if err != nil {
		return errors.New(err.Error())
	}
	ctx.Locals("data", dictItems)
	return nil
}
