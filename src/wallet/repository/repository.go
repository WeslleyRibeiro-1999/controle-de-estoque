package repository

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"gorm.io/gorm"
)

type repositoryWallet struct {
	db *gorm.DB
}

type RepositoryWallet interface {
	CreateRepository(wallet *models.Wallet) (*models.Wallet, error)
	UpdateWallet(wallet *models.Wallet) (*models.Wallet, error)
	GetWallet(id int32) (*models.Wallet, error)
}

var _ RepositoryWallet = (*repositoryWallet)(nil)

func NewRepositoryWallet(db *gorm.DB) RepositoryWallet {
	return &repositoryWallet{
		db: db,
	}
}

func (r *repositoryWallet) CreateRepository(wallet *models.Wallet) (*models.Wallet, error) {
	if err := r.db.Create(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *repositoryWallet) UpdateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	if err := r.db.Model(&models.Wallet{}).Where("id = ?", wallet.ID).Updates(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *repositoryWallet) GetWallet(id int32) (*models.Wallet, error) {
	var wallet *models.Wallet
	if err := r.db.First(&wallet, id).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}
