package households

import (
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, household models.Household) error {
	return db.Create(&household).Error
}
