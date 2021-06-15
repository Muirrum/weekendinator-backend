package db

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Base
	Token     string `gorm:"primarykey"`
	ExpiresAt time.Time
	UserID    uuid.UUID
}
