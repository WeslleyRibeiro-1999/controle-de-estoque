package repository

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Create(fornecedor *models.Fornecedor) (*models.Fornecedor, error)
	FindAll() (*[]models.Fornecedor, error)
	FindOne(fornecedor *models.Fornecedor) (*models.Fornecedor, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(fornecedor *models.Fornecedor) (*models.Fornecedor, error) {
	if err := r.db.Create(&fornecedor).Error; err != nil {
		return nil, err
	}

	return fornecedor, nil
}

func (r *repository) FindAll() (*[]models.Fornecedor, error) {
	var fornecedor []models.Fornecedor
	if err := r.db.Find(&fornecedor).Error; err != nil {
		return nil, err
	}

	return &fornecedor, nil
}

func (r *repository) FindOne(fornecedor *models.Fornecedor) (*models.Fornecedor, error) {
	if err := r.db.First(fornecedor, fornecedor.ID).Error; err != nil {
		return nil, err
	}

	return fornecedor, nil
}
