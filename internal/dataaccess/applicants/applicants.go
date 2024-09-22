package applicants

import (
	"github.com/mongj/gds-onecv-swe-assignment/internal/models"
	"gorm.io/gorm"
)

func List(db *gorm.DB) ([]models.Applicant, error) {
	var applicants []models.Applicant
	if err := db.Preload("Person").Find(&applicants).Error; err != nil {
		return nil, err
	}
	return applicants, nil
}

func Create(db *gorm.DB, applicant models.Person, relatives []models.Relative) error {
	// Create the person
	if err := db.Create(&applicant).Error; err != nil {
		return err
	}

	// Create the applicant
	a := models.Applicant{
		PersonID: applicant.ID,
		Person:   &applicant,
	}
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	// Create the relatives
	for _, r := range relatives {
		h := models.Household{
			PersonID: applicant.ID,
			Relative: r.Person,
			Relation: r.Relation,
		}
		if err := db.Create(&h).Error; err != nil {
			return err
		}
	}
	return nil
}
