package stat

import (
	"api/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type Repository struct {
	Database *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{Database: db}
}

func (repo *Repository) AddDirection(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Database.Find(&stat, "link_id = ? and directions_date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			LinkId:          linkId,
			DirectionsCount: 1,
			DirectionsDate:  currentDate,
		})
	} else {
		stat.DirectionsCount++
		repo.Database.Save(&stat)
	}
}
