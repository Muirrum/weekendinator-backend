package db

import (
	"log"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	Base
	Username     string
	Email        string
	FirstName    string
	LastName     string
	PasswordHash string
}

// Set UUID instead of numeric ID
func (base *User) BeforeCreate(tx *gorm.DB) error {
	base.Base.ID = uuid.New()
	log.Println("%s", base.Base.ID)

	return nil
}
