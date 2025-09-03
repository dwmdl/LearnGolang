package link

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Url       string         `gorm:"column:url" json:"url"`
	Hash      string         `gorm:"column:hash;uniqueIndex" json:"hash"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: RandStringRunes(10),
	}
}

var availableRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	randRunes := make([]rune, n)
	for i := range randRunes {
		randRunes[i] = availableRunes[rand.Intn(len(availableRunes))]
	}

	return string(randRunes)
}
