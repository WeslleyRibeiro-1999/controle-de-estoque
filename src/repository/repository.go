package repository

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Create(product *models.Produto) (*models.Produto, error)
	FindAll() (*[]models.Produto, error)
	FindOne(product *models.Produto) (*models.Produto, error)
}

var _ Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product *models.Produto) (*models.Produto, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) FindAll() (*[]models.Produto, error) {
	var products []models.Produto
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *repository) FindOne(product *models.Produto) (*models.Produto, error) {
	if err := r.db.First(product, product.ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}
