package db

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Set UUID instead of numeric ID
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	log.Println("%s", base.ID)

	return nil
}
