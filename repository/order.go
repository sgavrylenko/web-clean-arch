package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sgavrylenko/web-clean-arch/model"
)

// OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

// NewOrderRepository --> returns new order repository
func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: DB(),
	}
}

func (db *orderRepository) OrderProduct(userID int, productID int, quantity int) error {
	return db.connection.Create(&model.Order{
		ProductID: uint(productID),
		UserID:    uint(userID),
		Quantity:  quantity,
	}).Error
}
