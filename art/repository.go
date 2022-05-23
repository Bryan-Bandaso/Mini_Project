package art

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Art, error)
	FindByID(ID int) (Art, error)
	Create(artist Art) (Art, error)
	Update(artist Art) (Art, error)
	Delete(artist Art) (Art, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Art, error) {
	var artist []Art

	err := r.db.Find(&artist).Error

	return artist, err
}

func (r *repository) FindByID(ID int) (Art, error) {
	var artist Art

	err := r.db.Find(&artist).Error

	return artist, err
}

func (r *repository) Create(artist Art) (Art, error) {
	err := r.db.Create(&artist).Error

	return artist, err
}

func (r *repository) Update(artist Art) (Art, error) {
	err := r.db.Save(&artist).Error

	return artist, err
}

func (r *repository) Delete(artist Art) (Art, error) {
	err := r.db.Delete(&artist).Error

	return artist, err
}
