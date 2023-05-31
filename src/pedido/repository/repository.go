package repository

import (
	"github.com/WeslleyRibeiro-1999/controle-de-estoque/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetOrderID(id int32) (*models.Pedido, error)
	NewOrder(order *models.Pedido) (*models.Pedido, error)
	GetAllOrders() (*[]models.Pedido, error)
	NewProductOrder(order *models.ProdutosPedido) (*models.ProdutosPedido, error)
	UpdateOrder(pedido *models.Pedido) (*models.Pedido, error)
}

var _ Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) NewOrder(order *models.Pedido) (*models.Pedido, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) NewProductOrder(order *models.ProdutosPedido) (*models.ProdutosPedido, error) {
	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) GetOrderID(id int32) (*models.Pedido, error) {
	var order *models.Pedido
	if err := r.db.First(order, id).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *repository) GetAllOrders() (*[]models.Pedido, error) {
	var orders []models.Pedido
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

func (r *repository) UpdateOrder(pedido *models.Pedido) (*models.Pedido, error) {
	if err := r.db.Model(&models.Pedido{}).Where("id = ?", pedido.ID).Updates(pedido).Error; err != nil {
		return nil, err
	}

	return pedido, nil
}
