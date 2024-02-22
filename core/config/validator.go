package core_config

import (
	core_handler "gin-api/core/handler"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func ConfigValidator() {
	logger := &log.Logger
	logger.Info().Str("logtype", "ConfigValidator").Msg("Configuring Validator...")

	VALIDATE = validator.New()
	VALIDATE.RegisterValidation("beforenow", core_handler.IsBeforeNow)

	logger.Info().Str("logtype", "ConfigValidator").Msg("Validator CONFIGURED!")
}
