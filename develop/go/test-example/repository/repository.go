package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/my/testcode/model"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*model.Book, error)
	Create(id uuid.UUID, name string) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (p *repository) Create(id uuid.UUID, name string) error {
	book := &model.Book{
		ID:   id,
		Name: name,
	}

	return p.DB.Create(book).Error
}

func (p *repository) Get(id uuid.UUID) (*model.Book, error) {
	book := new(model.Book)

	err := p.DB.Where("id = $1", id).Find(book).Error

	return book, err
}
