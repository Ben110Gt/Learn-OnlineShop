package service

import (
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/util"
	"errors"
)

type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductByID(productID string) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	DeleteProduct(productID string) error
	UpdateProduct(productID string, updatedProduct *models.Product) (*models.Product, error)
}

type productService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *productService {
	return &productService{repo: repo}
}

// CreateProduct creates a new product
func (s *productService) CreateProduct(product *models.Product) (*models.Product, error) {
	exiting, err := s.repo.GetProductByID(product.ProductID)
	if err == nil && exiting != nil {
		return nil, errors.New("product name already in use")
	}
	newID, _ := util.GenerateID("P", configs.GetDB(), &models.Product{}, "product_id")
	product.ProductID = newID
	createdProduct, err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil

}

// GetProductByID retrieves a product by its ID
func (s *productService) GetProductByID(productID string) (*models.Product, error) {
	return s.repo.GetProductByID(productID)
}

// GetAllProducts retrieves all products
func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

// DeleteProduct deletes a product by its ID
func (s *productService) DeleteProduct(productID string) error {
	return s.repo.DeleteProduct(productID)
}

// Update product
func (s *productService) UpdateProduct(productID string, updatedProduct *models.Product) (*models.Product, error) {
	Product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	//update info
	if updatedProduct.ProductName != "" {
		Product.ProductName = updatedProduct.ProductName
	}
	if updatedProduct.Description != "" {
		Product.Description = updatedProduct.Description
	}
	if updatedProduct.Price != 0 {
		Product.Price = updatedProduct.Price
	}
	if updatedProduct.CategoryID != "" {
		Product.CategoryID = updatedProduct.CategoryID
	}
	if updatedProduct.ProductImage != "" {
		Product.ProductImage = updatedProduct.ProductImage
	}
	if updatedProduct.Stock != 0 {
		Product.Stock = updatedProduct.Stock
	}
	updatedProduct, err = s.repo.UpdateProduct(Product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}
