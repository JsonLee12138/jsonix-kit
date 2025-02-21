package service

import (
	"errors"
	"fmt"
	"json-server-kit/apps/auth/model/dto"
	"json-server-kit/apps/auth/model/enum"
	"json-server-kit/apps/auth/repository"
	"json-server-kit/apps/common/entity"
	"json-server-kit/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
	"gorm.io/gorm"
)

type LoginService struct {
	userRepository *repository.UserRepository
	uaparser       *uaparser.Parser
}

func NewLoginService(userRepository *repository.UserRepository, uparser *uaparser.Parser) *LoginService {
	return &LoginService{
		userRepository: userRepository,
		uaparser:       uparser,
	}
}

func (service *LoginService) HelloWord() string {
	return "Hello Word!"
}

func (service *LoginService) GetUserInfoByUsername(username string) (*entity.User, error) {
	user, err := service.userRepository.GetUserInfoByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("password_error") // 账号不存在
		}
		return nil, err
	}
	return user, nil
}

func (service *LoginService) UsernameLoginWithCaptcha(c *fiber.Ctx, loginDTO dto.UsernameLoginWithCaptchaDTO) (*entity.User, error) {
	user, err := service.userRepository.GetUserInfoByUsername(loginDTO.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("password_error") // 账号不存在
		}
		return nil, err
	}
	if !user.Status {
		return nil, errors.New("user_disabled")
	}

	// 验证密码
	if !utils.CompareHashAndPassword(user.Password, loginDTO.Password) {
		// redis 记录错误次数(5分钟内错误6次永久封禁)
		// 前期只能联系管理员解封(后期通过手机验证码或邮箱解封)
		count, err := service.userRepository.SetPasswordError(user.ID)
		if err != nil {
			return nil, err
		}
		if count >= 6 {
			err := service.userRepository.DisableUserForever(user.ID)
			if err != nil {
				return nil, err
			}
			service.userRepository.ClearPasswordError(user.ID)
		}
		return nil, errors.New("password_error")
	}
	userAgent := c.Get("User-Agent")

	parser, err := uaparser.New("./config/regexes.yaml")
	if err != nil {
		return nil, err
	}

	deviceFamily := parser.Parse(userAgent).Device.Family
	if strings.ToLower(deviceFamily) == "spider" || strings.ToLower(deviceFamily) == "Bot" {
		return nil, errors.New("automated_request")
	}
	// 如果需要分设备类型单独登录，请放开以下代码
	// var deviceType enum.DeviceType
	// if deviceFamily == "iPhone" || deviceFamily == "Android" || deviceFamily == "iPad" {
	// 	deviceType = enum.DeviceTypeMobile
	// } else {
	// 	deviceType = enum.DeviceTypeDesktop
	// }
	// 如果要分web和app登录，请放开以下代码
	// 如果是web登录需要设置cookie, 否则headers设置双token
	var clientType enum.ClientType
	if enum.IsApp(userAgent) {
		clientType = enum.ClientTypeApp
	} else {
		clientType = enum.ClientTypeWeb
	}
	fmt.Println(clientType)
	if clientType == enum.ClientTypeApp {

	}
	return user, nil
}

func (service *LoginService) GetUAParser(userAgent string) *uaparser.Client {
	return service.uaparser.Parse(userAgent)
}

func (service *LoginService) UserEnabled(user *entity.User) error {
	if !user.Status {
		return errors.New("user_disabled")
	}
	return nil
}

func (service *LoginService) CheckPassword(user *entity.User, password string) error {
	if !utils.CompareHashAndPassword(user.Password, password) {
		// redis 记录错误次数(5分钟内错误6次永久封禁)
		// 前期只能联系管理员解封(后期通过手机验证码或邮箱解封)
		count, err := service.userRepository.SetPasswordError(user.ID)
		if err != nil {
			return err
		}
		if count >= 6 {
			err := service.userRepository.DisableUserForever(user.ID)
			if err != nil {
				return err
			}
			service.userRepository.ClearPasswordError(user.ID)
		}
		return errors.New("password_error")
	}
	return nil
}

// func (service *LoginService) GenerateToken(c *fiber.Ctx, loginDTO dto.UsernameLoginWithCaptchaDTO) (string, error) {
// }
