package repositories

import (
	"OnlineShop/Api/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

// CreateCategory creates a new category in the database
func (r *CategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	if err := r.DB.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

// GetCategoryByName retrieves a category by its name
func (r *CategoryRepository) GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	if err := r.DB.Where("category_name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAllCategories retrieves all categories from the database
func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Delete
func (r *CategoryRepository) DeleteCategories(categoryID string) error {
	if err := r.DB.Where("category_id = ?", categoryID).Unscoped().Delete(&models.Category{}).Error; err != nil {
		return err
	}
	return nil
}
