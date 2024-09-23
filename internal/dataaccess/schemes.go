package dataaccess

import (
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"gorm.io/gorm"
)

func ListSchemes(db *gorm.DB) ([]models.Scheme, error) {
	var schemes []models.Scheme
	if err := db.Preload("Benefits").Preload("Criteria").Find(&schemes).Error; err != nil {
		return nil, err
	}
	return schemes, nil
}
