package repositories

import (
	"OnlineShop/Api/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database
func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := r.DB.Create(product).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Preload("Category").Where("product_id = ?", product.ProductID).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductByID retrieves a product by its ID
func (r *ProductRepository) GetProductByID(productID string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Preload("Category").
		Where("product_id = ?", productID).
		First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductByName retrieves a product by its name
func (r *ProductRepository) GetProductByName(name string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Preload("Category").
		Where("product_name = ?", name).
		First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProducts retrieves all products from the database
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct updates an existing product in the database
func (r *ProductRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := r.DB.Save(product).Error; err != nil {
		return nil, err
	}
	// โหลดข้อมูล Category หลังอัปเดต
	if err := r.DB.Preload("Category").Where("product_id = ?", product.ProductID).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct deletes a product from the database by its ID
func (r *ProductRepository) DeleteProduct(productID string) error {
	if err := r.DB.Where("product_id = ?", productID).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}

// Update status from admin
func (r *ProductRepository) UpdateStock(productID string, quantity int) error {
	return r.DB.Model(&models.Product{}).
		Where("product_id = ? AND stock >= ?", productID, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

// U
