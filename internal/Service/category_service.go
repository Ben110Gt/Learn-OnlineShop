package service

import (
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/util"
	"errors"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	exiting, err := s.repo.GetCategoryByName(category.CategoryName)
	if err == nil && exiting != nil {
		return nil, errors.New("category name already in use")
	}
	newID, _ := util.GenerateID("C", configs.GetDB(), &models.Category{}, "category_id")
	category.CategoryID = newID
	createdCategory, err := s.repo.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

// GetAllCategories retrieves all categories
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) DeleteCategories(categoryID string) error {
	return s.repo.DeleteCategories(categoryID)
}
