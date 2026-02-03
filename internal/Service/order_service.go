package service

import (
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/util"
	"errors"
	"fmt"
)

type OrderService struct {
	cartRepo  *repositories.CartRepository
	orderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, cartRepo *repositories.CartRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, cartRepo: cartRepo}
}

func (s *OrderService) CreateOrder(userID string) (*models.Order, error) {
	carts, err := s.cartRepo.GetCartByUserID(userID)

	if err != nil {
		return nil, err
	}
	fmt.Printf("DEBUG: userID=%s, carts=%+v\n", userID, carts)
	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}
	newOrderID, _ := util.GenerateID("Order", configs.GetDB(), &models.Order{}, "order_id")
	order := &models.Order{
		OrderID: newOrderID,
		UserID:  userID,
		Status:  "pending",
		Total:   0.0,
		Items:   []models.OrderItem{},
	}

	for _, cart := range carts {
		newOrderItemID, _ := util.GenerateID("OrderItem", configs.GetDB(), &models.OrderItem{}, "order_item_id")
		orderItem := models.OrderItem{
			OrderItemID: newOrderItemID,
			OrderID:     order.OrderID,
			ProductID:   cart.ProductID,
			Quantity:    float64(cart.Quantity),
			Price:       cart.Product.Price,
		}
		order.Items = append(order.Items, orderItem)
		order.Total += float64(cart.Product.Price) * float64(cart.Quantity)
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	if err := s.cartRepo.DeleteByUser(userID); err != nil {
		return nil, err
	}

	return order, nil
}
func (s *OrderService) ConfirmOrder(orderID string) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status != "pending" {
		return errors.New("order is not in pending status")
	}

	if err := s.orderRepo.UpdateOrderStatus(orderID, "confirmed"); err != nil {
		return err
	}

	for _, item := range order.Items {
		if err := s.orderRepo.DeductStock(item.ProductID, int(item.Quantity)); err != nil {
			return err
		}
	}

	return nil
}
func (s *OrderService) GetOrderByID(userID string) (*models.Order, error) {
	return s.orderRepo.GetOrderByID(userID)
}
func (s *OrderService) GetAllOrder() ([]models.Order, error) {
	return s.orderRepo.GetAllOrder()
}
