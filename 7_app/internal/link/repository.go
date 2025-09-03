package link

import "api/pkg/db"

type Repository struct {
	Database *db.DB
}

func NewRepository(database *db.DB) *Repository {
	return &Repository{Database: database}
}

func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.Database.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}
