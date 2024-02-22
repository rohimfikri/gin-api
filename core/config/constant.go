package core_config

import (
	core_type "gin-api/core/type"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var ENV *core_type.Config
var DB_SYS *gorm.DB
var VALIDATE *validator.Validate
