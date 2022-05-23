package artist

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Artist, error)
	FindByID(ID int) (Artist, error)
	Create(artist Artist) (Artist, error)
	Update(artist Artist) (Artist, error)
	Delete(artist Artist) (Artist, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Artist, error) {
	var artist []Artist

	err := r.db.Find(&artist).Error

	return artist, err
}

func (r *repository) FindByID(ID int) (Artist, error) {
	var artist Artist

	err := r.db.Find(&artist).Error

	return artist, err
}

func (r *repository) Create(artist Artist) (Artist, error) {
	err := r.db.Create(&artist).Error

	return artist, err
}

func (r *repository) Update(artist Artist) (Artist, error) {
	err := r.db.Save(&artist).Error

	return artist, err
}

func (r *repository) Delete(artist Artist) (Artist, error) {
	err := r.db.Delete(&artist).Error

	return artist, err
}
