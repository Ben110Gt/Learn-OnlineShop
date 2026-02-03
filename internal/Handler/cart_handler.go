package handler

import (
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	CartService service.CartService
}

func NewCartHandler(CartService service.CartService) *CartHandler {
	return &CartHandler{CartService: CartService}
}
func (cc *CartHandler) AddCart(c *fiber.Ctx) error {
	CartItem := new(models.Cart)
	if err := c.BodyParser(CartItem); err != nil {
		c.Status(400).SendString(err.Error())
	}
	createCart, err := cc.CartService.AddToCart(CartItem)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	res := models.CartResponse{
		CartID:      createCart.CartID,
		UserID:      createCart.UserID,
		ProductID:   createCart.ProductID,
		ProductName: createCart.Product.ProductName,
		// Price: createCart.,
		Quantity: createCart.Quantity,
	}
	return c.Status(200).JSON(res)

}
func (cc *CartHandler) GetCartbyID(c *fiber.Ctx) error {
	CartID := c.Params("user_id")
	carts, err := cc.CartService.GetCartByUserID(CartID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}
	var response []models.CartResponse
	for _, c := range carts {
		resp := models.CartResponse{
			CartID:      c.CartID,
			UserID:      c.UserID,
			ProductID:   c.Product.ProductID,
			ProductName: c.Product.ProductName,
			// Price:       c.Product.Price,
			Quantity: c.Quantity,
		}
		response = append(response, resp)

	}
	return c.Status(200).JSON(response)
}
