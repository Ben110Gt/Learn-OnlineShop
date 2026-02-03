package service

import (
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/util"
)

type CartService struct {
	CartRepo *repositories.CartRepository
}

func NewCartService(cartRepo *repositories.CartRepository) *CartService {
	return &CartService{CartRepo: cartRepo}
}

// AddToCart adds a product to the cart
func (s *CartService) AddToCart(cartItem *models.Cart) (*models.Cart, error) {
	newID, _ := util.GenerateID("Cart", configs.GetDB(), &models.Cart{}, "cart_id")
	cartItem.CartID = newID
	addCart, err := s.CartRepo.AddToCart(cartItem)
	if err != nil {
		return nil, err
	}
	return addCart, nil
}

// GetCartByUserID retrieves cart items by user ID
func (s *CartService) GetCartByUserID(userID string) ([]models.Cart, error) {
	return s.CartRepo.GetCartByUserID(userID)
}
