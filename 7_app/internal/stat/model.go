package stat

import (
	"time"

	"gorm.io/datatypes"
)

type Stat struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	LinkId          uint           `gorm:"column:link_id" json:"link_id"`
	DirectionsCount uint           `gorm:"column:directions_count" json:"directions_count"`
	DirectionsDate  datatypes.Date `gorm:"column:directions_date" json:"directions_date"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"-"`
}
