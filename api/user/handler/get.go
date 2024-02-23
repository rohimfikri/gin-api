package user_handler

import (
	"fmt"
	user_type "gin-api/api/user/type"
	core_config "gin-api/core/config"
	core_model "gin-api/core/model"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func GetActiveUsers(param *user_type.GetActiveUsersRequest) (*[]map[string]interface{}, error) {
	logger := &log.Logger

	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(param)
	if validateErr != nil {
		logger.Info().Str("logtype", "GetActiveUsers").Err(validateErr)
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	var users []map[string]interface{}
	var like_cond []string
	like_cond = append(like_cond, "status=@status")

	like_value := make(map[string]interface{})
	like_value["status"] = 1

	if param.EmailLike != "" {
		like_cond = append(like_cond, "email like @email")
		like_value["email"] = fmt.Sprintf("%%%v%%", param.EmailLike)
	}
	// fmt.Println(like_cond, like_value)

	if err := core_model.SearchActiveUsers(core_config.DB_SYS, &like_cond, &like_value, &users); err != nil {
		msg := "Failed retrieve data"
		logger.Info().Str("logtype", "GetActiveUsers").Msg(msg)
		return nil, fmt.Errorf("%v", msg)
	}

	return &users, nil
}

func GetUserByID(param *user_type.GetUserByIDRequest) (*core_model.User, error) {
	logger := &log.Logger

	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(param)
	if validateErr != nil {
		logger.Info().Str("logtype", "GetUserByID").Err(validateErr)
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	u := core_model.User{}

	if err := u.FindByID(core_config.DB_SYS, &param.ID); err != nil || u.ID == "" {
		msg := "User Not Found"
		logger.Info().Str("logtype", "GetUserByID").Msg(msg)
		return nil, fmt.Errorf("%v", msg)
	}

	return &u, nil
}
