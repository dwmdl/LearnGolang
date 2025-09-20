package link

import (
	"api/pkg/db"

	"gorm.io/gorm/clause"
)

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

func (repo *Repository) GetByHash(hash string) (link *Link, err error) {
	result := repo.Database.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *Repository) Update(link *Link) (*Link, error) {
	result := repo.Database.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *Repository) Delete(id uint64) error {
	result := repo.Database.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *Repository) GetById(id uint64) (*Link, error) {
	var link Link

	result := repo.Database.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}
