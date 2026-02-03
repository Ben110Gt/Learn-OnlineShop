package repositories

import (
	"OnlineShop/Api/internal/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

// AddToCart adds a product to the cart
func (r *CartRepository) AddToCart(cartItem *models.Cart) (*models.Cart, error) {
	if err := r.DB.Create(cartItem).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Preload("Product.Category").Where("cart_id= ?", cartItem.CartID).First(cartItem).Error; err != nil {
		return nil, err
	}
	return cartItem, nil
}

// GetCartByUserID retrieves cart items by user ID
func (r *CartRepository) GetCartByUserID(userID string) ([]models.Cart, error) {
	var cartItems []models.Cart
	if err := r.DB.Preload("Product.Category").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

// Delete
func (r *CartRepository) DeleteByUser(userID string) error {
	return r.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
