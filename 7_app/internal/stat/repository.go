package stat

import (
	"api/pkg/db"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
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

func (repo *Repository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var query string

	switch by {
	case GroupByDay:
		query = "to_char(directions_date, 'YYYY-MM-DD') as period, sum(directions_count) as directions_sum"
	case GroupByMonth:
		query = "to_char(directions_date, 'YYYY-MM') as period, sum(directions_count) as directions_sum"
	}

	sessionQuery := repo.Database.Table("stats").
		Select(query).
		Session(&gorm.Session{})

	sessionQuery.Where("directions_date between ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
