package handler

import (
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(CategoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: CategoryService}
}

func (g *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		c.Status(400).SendString(err.Error())
	}
	createCat, err := g.CategoryService.CreateCategory(category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(createCat)

}
func (g *CategoryHandler) GetAllCategory(c *fiber.Ctx) error {
	category, err := g.CategoryService.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch Product",
		})
	}
	return c.Status(200).JSON(category)

}
func (g *CategoryHandler) DeleteCategoryByid(c *fiber.Ctx) error {
	Catid := c.Params("category_id")
	err := g.CategoryService.DeleteCategories(Catid)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"message": "category deleted successfully"})
}
