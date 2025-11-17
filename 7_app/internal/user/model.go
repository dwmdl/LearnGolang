package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"-"`
	Email     string         `gorm:"column:email;index'" json:"email"`
	Name      string         `gorm:"column:name" json:"name"`
	Password  string         `gorm:"column:password" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}
