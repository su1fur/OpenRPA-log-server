package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"openRPA-log-server/app/models"
)

func ValidateExistTaskLine(fl validator.FieldLevel) bool {
	task := models.Task{}
	taskLineID := fl.Field().String()
	if err := models.DB.First(&task, "task_line_id = ?", taskLineID).Error; err != nil {
		return false
	}

	return true
}
