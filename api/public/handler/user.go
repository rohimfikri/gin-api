package public_handler

import (
	"fmt"
	public_type "gin-api/api/public/type"
	core_config "gin-api/core/config"
	core_model "gin-api/core/model"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func RegisterUser(user *public_type.UserRegisterRequest) (interface{}, error) {
	logger := &log.Logger

	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		logger.Info().Str("logtype", "RegisterUser").Err(validateErr)
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	u := core_model.User{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  &user.LastName,
	}

	return u.SaveUser(core_config.DB_SYS)
}

func ChangePassword(user *public_type.ChangePasswordRequest) (interface{}, error) {
	logger := &log.Logger

	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		logger.Info().Str("logtype", "ChangePassword").Err(validateErr)
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	if user.SecretKey != core_config.ENV.JWT_SECRET_KEY {
		msg := "Invalid Key"
		logger.Info().Str("logtype", "ChangePassword").Msg(msg)
		return nil, fmt.Errorf("%v", msg)
	}

	// Prepare Data
	u := core_model.User{}

	if err := u.FindByUsername(core_config.DB_SYS, &user.Username); err != nil || u.ID == "" {
		msg := "User Not Found"
		logger.Info().Str("logtype", "ChangePassword").Msg(msg)
		return nil, fmt.Errorf("%v", msg)
	}

	return u.ChangePassword(core_config.DB_SYS)
}
