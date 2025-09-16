package tasks

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
	logger = config.GetLogger("tasksHandler")
	db = config.GetSQLite()
	validate = validator.New()
}
