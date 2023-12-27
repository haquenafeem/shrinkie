package repository

import (
	"github.com/haquenafeem/shrinkie/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&model.URL{})
}

func (r *Repository) CreateURL(url *model.URL) error {
	return r.db.Create(url).Error
}

func (r *Repository) GetUrl(random string) (*model.URL, error) {
	var url model.URL

	err := r.db.Where("random_string = ?", random).Find(&url).Error

	return &url, err
}

func (r *Repository) GetAll() ([]model.URL, error) {
	var urls []model.URL

	err := r.db.Find(&urls).Error

	return urls, err
}

func NewMust(db *gorm.DB) *Repository {
	repo, err := New(db)
	if err != nil {
		panic(err)
	}

	return repo
}

func New(db *gorm.DB) (*Repository, error) {
	repo := &Repository{
		db: db,
	}

	err := repo.Migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}
