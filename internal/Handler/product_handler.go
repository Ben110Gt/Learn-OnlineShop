package handler

import (
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/service"
	"OnlineShop/Api/internal/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// Handler CreateProduct
func (p *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	productName := c.FormValue("productname")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	categoryID := c.FormValue("categoryID")
	stockStr := c.FormValue("stock")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid price"})
	}

	stock := 0
	if stockStr != "" {
		s, err := strconv.Atoi(stockStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "stock must be a valid integer"})
		}
		stock = s
	}

	profilePath, err := util.UploadFileProduct(c, "Product_Image", "./Upload/Product/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to upload profile"})
	}

	product := &models.Product{
		ProductName:  productName,
		Description:  description,
		Price:        price,
		CategoryID:   categoryID,
		ProductImage: profilePath,
		Stock:        stock,
	}

	createdProduct, err := p.productService.CreateProduct(product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	res := models.ProductResponse{
		ProductID:     createdProduct.ProductID,
		ProductName:   createdProduct.ProductName,
		Description:   createdProduct.Description,
		Price:         createdProduct.Price,
		CategoryID:    createdProduct.CategoryID,
		Category:      createdProduct.Category.CategoryName,
		Product_Image: createdProduct.ProductImage,
		Stock:         createdProduct.Stock,
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "product created successfully",
		"data":    res,
	})
}

// Handler GetProduct By ID
func (p *ProductHandler) GetProductrByID(c *fiber.Ctx) error {
	PID := c.Params("product_id")
	product, err := p.productService.GetProductByID(PID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}
	product_res := models.ProductResponse{
		ProductID:     product.ProductID,
		ProductName:   product.ProductName,
		Description:   product.Description,
		Price:         product.Price,
		CategoryID:    product.CategoryID,
		Category:      product.Category.CategoryName,
		Product_Image: product.ProductImage,
		Stock:         product.Stock,
	}
	return c.Status(200).JSON(product_res)
}

// Get All Product in Routes Public
func (p *ProductHandler) GetAllProductPublic(c *fiber.Ctx) error {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch Product",
		})
	}
	var pres []models.ProductResponses
	for _, p := range products {
		pres = append(pres, models.ProductResponses{
			ProductID:     p.ProductID,
			ProductName:   p.ProductName,
			Description:   p.Description,
			Price:         p.Price,
			CategoryID:    p.CategoryID,
			Category:      p.Category.CategoryName,
			Product_Image: p.ProductImage,
		})
	}
	return c.Status(200).JSON(pres)
}

// Get AllProduct Admin
func (p *ProductHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch Product",
		})
	}
	var pres []models.ProductResponse
	for _, p := range products {
		pres = append(pres, models.ProductResponse{
			ProductID:     p.ProductID,
			ProductName:   p.ProductName,
			Description:   p.Description,
			Price:         p.Price,
			CategoryID:    p.CategoryID,
			Category:      p.Category.CategoryName,
			Product_Image: p.ProductImage,
			Stock:         p.Stock,
		})
	}
	return c.Status(200).JSON(pres)
}

// Delect Product
func (p *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	PID := c.Params("product_id")
	err := p.productService.DeleteProduct(PID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "user deleted successfully",
	})

}

// Update Product
func (p *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	PID := c.Params("product_id")

	productName := c.FormValue("productname")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	categoryID := c.FormValue("categoryID")
	stockStr := c.FormValue("stock")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid price"})
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid stock"})
	}

	profilePath, err := util.UploadFileProduct(c, "Product_Image", "./Upload/Product/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to upload product image"})
	}
	UpdateProduct := &models.Product{
		ProductName:  productName,
		Description:  description,
		Price:        price,
		CategoryID:   categoryID,
		ProductImage: profilePath,
		Stock:        stock,
	}
	product, err := p.productService.UpdateProduct(PID, UpdateProduct)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	res := models.ProductResponse{
		ProductID:     product.ProductID,
		ProductName:   product.ProductName,
		Description:   product.Description,
		Price:         product.Price,
		CategoryID:    product.CategoryID,
		Category:      product.Category.CategoryName,
		Product_Image: product.ProductImage,
		Stock:         product.Stock,
	}
	return c.Status(200).JSON(res)

}
