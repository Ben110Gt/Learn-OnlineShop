package service

import (
	"OnlineShop/Api/internal/configs"
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/repositories"
	"OnlineShop/Api/internal/util"
	"errors"
)

type UserService interface {
	RegisterUser(user *models.User) (*models.User, error)
	RegisterAdmin(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUser(userID string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	Login(phone, password string) (models.LoginResponse, error)
	UpdateUser(userID string, updatedUser *models.User) (*models.User, error)
	DeleteUser(userID string) error
}
type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo: repo}
}

// Service to register user
func (s *userService) RegisterUser(user *models.User) (*models.User, error) {
	exiting, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && exiting != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword := util.HashPassword(user.Password)
	newID, _ := util.GenerateID("U", configs.GetDB(), &models.User{}, "user_id")
	user = &models.User{
		UserID:       newID,
		Username:     user.Username,
		Email:        user.Email,
		Password:     hashedPassword,
		Phone:        user.Phone,
		Role:         "user",
		ProfileImage: user.ProfileImage,
	}
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

// Service to register user
func (s *userService) RegisterAdmin(user *models.User) (*models.User, error) {
	exiting, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && exiting != nil {
		return nil, errors.New("email already in use")
	}
	hashedPassword := util.HashPassword(user.Password)
	newID, _ := util.GenerateID("U", configs.GetDB(), &models.User{}, "user_id")
	user = &models.User{
		UserID:       newID,
		Username:     user.Username,
		Email:        user.Email,
		Password:     hashedPassword,
		Phone:        user.Phone,
		Role:         user.Role,
		ProfileImage: user.ProfileImage,
	}
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

// Service to get user by email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

// Service to get user by userID
func (s *userService) GetUser(userID string) (*models.User, error) {
	return s.repo.GetUser(userID)
}

// Service to get user by userID
func (s *userService) GetAllUser() ([]models.User, error) {
	return s.repo.GetAllUser()
}

// Service to login user
func (s *userService) Login(phone, password string) (models.LoginResponse, error) {
	user, err := s.repo.GetUserByPhone(phone)
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid phone or password")
	}
	if !util.CheckPassword(user.Password, password) {
		return models.LoginResponse{}, errors.New("invalid phone or password")
	}
	token, err := util.GenerateJWT(user.UserID, user.Phone, user.Role)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return models.LoginResponse{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}, nil
}

func (s *userService) UpdateUser(userID string, updatedUser *models.User) (*models.User, error) {
	Users, err := s.repo.GetUser(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// อัปเดตข้อมูลที่ได้รับมา
	if updatedUser.Username != "" {
		Users.Username = updatedUser.Username
	}
	if updatedUser.Email != "" {
		Users.Email = updatedUser.Email
	}
	if updatedUser.Phone != "" {
		Users.Phone = updatedUser.Phone
	}
	if updatedUser.ProfileImage != "" {
		Users.ProfileImage = updatedUser.ProfileImage
	}
	if updatedUser.Password != "" {
		hashedPassword := util.HashPassword(updatedUser.Password)
		Users.Password = string(hashedPassword)
	}
	updatedUser, err = s.repo.UpdateUser(Users)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// Service to delete user by userID
func (s *userService) DeleteUser(userID string) error {
	return s.repo.DeleteUser(userID)
}
