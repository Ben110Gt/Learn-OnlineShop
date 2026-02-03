package repositories

import (
	"OnlineShop/Api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
	GetUser(userID string) (*models.User, error)
	DeleteUser(userID string) error
	UpdateUser(user *models.User) (*models.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Query to create user
func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Query to get user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Query to get all user
func (r *userRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Query to get user by Phone
func (r *userRepository) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("phone= ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Query to get user
func (r *userRepository) GetUser(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Query to delete user by user_id
func (r *userRepository) DeleteUser(userID string) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

// Query UpdateUser updates an existing user in the database
func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
