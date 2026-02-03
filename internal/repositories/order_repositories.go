package repositories

import (
	"OnlineShop/Api/internal/models"
	"errors"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}
func (r *OrderRepository) CreateOrder(order *models.Order) error {

	return r.DB.Create(order).Error

}

func (r *OrderRepository) UpdateOrderStatus(orderID, status string) error {
	return r.DB.Model(&models.Order{}).Where("order_id = ?", orderID).Update("status", status).Error
}

func (r *OrderRepository) DeductStock(productID string, quantity int) error {
	var product models.Product
	if err := r.DB.First(&product, "product_id = ?", productID).Error; err != nil {
		return err
	}
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}
	return r.DB.Model(&product).Update("stock", product.Stock-quantity).Error
}

func (r *OrderRepository) GetOrderByID(orderID string) (*models.Order, error) {
	var order models.Order
	err := r.DB.Preload("Items.Product").First(&order, "order_id = ?", orderID).Error
	return &order, err
}
func (r *OrderRepository) GetAllOrder() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("Items.Product").Find(&orders).Error
	return orders, err
}
