package model

import (
	"time"
)

type User struct {
	// gorm.Model
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ID        string    `gorm:"primaryKey;size:30"`
	Username  string    `gorm:"unique;size:20;not null"`
	Password  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	FirstName string    `gorm:"size:50;not null"`
	LastName  *string   `gorm:"size:50"`
	Status    uint      `gorm:"default:1;precision:1;size:1;not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}
