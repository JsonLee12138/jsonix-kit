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
// @Param dictTypeQuery query dto.DictTypeQuery true "字典查询条件"
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

// @Summary 新增字典类型
// @Description 新增字典类型
// @Tags common
// @Accept json
// @Produce json
// @Param createDictTypeDTO body dto.CreateDictTypeDTO true "新增字典类型信息"
// @Success 200 {object} utils.BaseResponseVO{data=entity.DictType}
// @Router /common/dict/types [post]
func (c *DictController) CreateDictType(ctx *fiber.Ctx) error {
	var body dto.CreateDictTypeDTO
	if err := ctx.BodyParser(&body); err != nil {
		return utils.NewBadParameters()
	}
	res, err := c.DictService.CreateDictType(body)
	if err != nil {
		return errors.New(err.Error())
	}
	ctx.Locals("data", res)
	return nil
}

// @Summary 获取字典值列表
// @Description 获取字典值列表
// @Tags common
// @Accept json
// @Produce json
// @Param code path string true "字典类型编码" Enums(sys_user_sex,sys_user_status)
// @Success 200 {object} utils.BaseResponseVO{data=[]entity.DictItem}
// @Router /common/dict/${code}/values [get]
func (c *DictController) GetDictValues(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
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
