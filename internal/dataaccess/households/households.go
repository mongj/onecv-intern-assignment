package households

import (
	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, household models.Household) error {
	return db.Create(&household).Error
}

func ReadByPersonID(db *gorm.DB, personID uuid.UUID) ([]models.Household, error) {
	var households []models.Household
	if err := db.Preload("Relative").Where("person_id = ?", personID).Find(&households).Error; err != nil {
		return []models.Household{}, err
	}
	return households, nil
}
