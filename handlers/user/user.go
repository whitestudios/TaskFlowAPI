package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/whitestudios/TaskFlowAPI/config"
	"gorm.io/gorm"
)

var (
	logger   *config.Logger
	db       *gorm.DB
	validate *validator.Validate
)

func Init() {
	validate = validator.New()
	logger = config.GetLogger("userHandler")
	db = config.GetSQLite()
}
