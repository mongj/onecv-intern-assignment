package people

import (
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, person models.Person) (models.Person, error) {
	res := db.Create(&person).Error
	return person, res
}
