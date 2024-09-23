package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mongj/gds-onecv-swe-assignment/internal/enums"
	"gorm.io/gorm"
)

type Application struct {
	ID                uuid.UUID               `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ApplicantID       uuid.UUID               `gorm:"type:uuid;not null"`
	Applicant         *Applicant              `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SchemeID          uuid.UUID               `gorm:"type:uuid;not null"`
	Scheme            *Scheme                 `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ApplicationStatus enums.ApplicationStatus `gorm:"type:application_status;not null"`
}

func (a *Application) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	// Check that applicant id exists and is valid
	applicant, err := ReadApplicant(tx, a.ApplicantID)
	if applicant == nil {
		return fmt.Errorf("applicant with ID %s does not exist", a.ApplicantID)
	}
	if err != nil {
		return err
	}

	// Check that the scheme ID exists and is valid
	scheme, err := ReadScheme(tx, a.SchemeID)
	if scheme == nil {
		return fmt.Errorf("scheme with ID %s does not exist", a.SchemeID)
	}
	if err != nil {
		return err
	}

	// Check that the applicant does not have another pending application for the same scheme
	var count int64
	if err := tx.Model(&Application{}).
		Where("applicant_id = ? AND scheme_id = ? AND application_status = ?", a.ApplicantID, a.SchemeID, enums.ApplicationStatusPending).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("applicant with ID %s already has a pending application for scheme with ID %s", a.ApplicantID, a.SchemeID)
	}

	return nil
}

// CreateApplications creates multiple applications in a single transaction
// If any of the applications fail to be created, the transaction is rolled back
func CreateApplications(db *gorm.DB, appl []Application) error {
	tx := db.Begin()
	// Create the applications
	for _, a := range appl {
		if err := a.Create(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func ListApplications(db *gorm.DB) ([]Application, error) {
	var a []Application
	if err := db.Preload("Applicant").Preload("Scheme").Find(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}
