package handler

import (
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

// User Create Order
func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	req := new(models.CreateOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	createdOrder, err := o.OrderService.CreateOrder(req.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var items []models.OrderItemResponse
	for _, item := range createdOrder.Items {
		items = append(items, models.OrderItemResponse{
			OrderItemID: item.OrderItemID,
			ProductID:   item.ProductID,
			ProductName: item.Product.ProductName,
			Quantity:    int(item.Quantity),
			Price:       int(item.Price),
		})
	}

	orderRespon := models.OrderResponse{
		OrderID: createdOrder.OrderID,
		UserID:  createdOrder.UserID,
		Total:   int(createdOrder.Total),
		Status:  createdOrder.Status,
		Items:   items,
	}
	return c.Status(200).JSON(orderRespon)
}

// Confirm By Admin
func (o *OrderHandler) ConfirmOrder(c *fiber.Ctx) error {
	OID := c.Params("order_id")
	err := o.OrderService.ConfirmOrder(OID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Order Confirm"})

}

// Get All Order
func (o *OrderHandler) GetAllOrder(c *fiber.Ctx) error {
	order, err := o.OrderService.GetAllOrder()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch Product",
		})
	}

	var orders []models.OrderResponse
	for _, ord := range order {
		var items []models.OrderItemResponse
		for _, item := range ord.Items {
			items = append(items, models.OrderItemResponse{
				OrderItemID: item.OrderItemID,
				ProductID:   item.ProductID,
				ProductName: item.Product.ProductName,
				Quantity:    int(item.Quantity),
				Price:       int(item.Price),
			})
		}
		orders = append(orders, models.OrderResponse{
			OrderID: ord.OrderID,
			UserID:  ord.UserID,
			Total:   int(ord.Total),
			Status:  ord.Status,
			Items:   items,
		})
	}
	return c.Status(200).JSON(orders)
}
