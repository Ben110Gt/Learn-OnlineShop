package handler

import (
	"OnlineShop/Api/internal/models"
	"OnlineShop/Api/internal/service"
	"OnlineShop/Api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Handler to register user
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	username := c.FormValue("Username")
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	phone := c.FormValue("Phone")

	profilePath, err := util.UploadFile(c, "Profile_User", "./Upload/User/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to upload profile"})
	}
	user := &models.User{
		Username:     username,
		Email:        email,
		Password:     password,
		Phone:        phone,
		Role:         "user",
		ProfileImage: profilePath,
	}
	createdUser, err := h.userService.RegisterUser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	createdUser.Password = ""
	return c.Status(201).JSON(createdUser)
}

// ...other handlers...

// Handler to register admin
func (h *UserHandler) RegisterAdmin(c *fiber.Ctx) error {
	username := c.FormValue("Username")
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	phone := c.FormValue("Phone")
	role := c.FormValue("Role")

	profilePath, err := util.UploadFile(c, "Profile_User", "./Upload/User/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to upload profile"})
	}
	user := &models.User{
		Username:     username,
		Email:        email,
		Password:     password,
		Phone:        phone,
		ProfileImage: profilePath,
		Role:         role,
	}
	createdUser, err := h.userService.RegisterUser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	createdUser.Password = ""
	return c.Status(201).JSON(createdUser)
}

// Handler to get user by email
func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	user.Password = ""
	return c.Status(200).JSON(user)
}

// Handler to get user by userID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	user, err := h.userService.GetUser(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	user.Password = ""
	return c.Status(200).JSON(user)
}

// ...other handlers...
// Handler to update user
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	// Find user by ID
	userID := c.Params("user_id")
	// Parse the form data
	username := c.FormValue("Username")
	email := c.FormValue("Email")
	phone := c.FormValue("Phone")
	profilePath, err := util.UploadFile(c, "Profile_User", "./Upload/User/")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to upload profile"})
	}
	updatedUser := &models.User{
		Username:     username,
		Email:        email,
		Phone:        phone,
		ProfileImage: profilePath,
	}
	user, err := h.userService.UpdateUser(userID, updatedUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user.Password = ""
	return c.Status(200).JSON(user)
}

// Handler to get all users
func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUser()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}
	for i := range users {
		users[i].Password = ""
	}
	return c.Status(200).JSON(users)
}

// Handler to login user
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	loginResponse, err := h.userService.Login(loginRequest.Phone, loginRequest.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(loginResponse)
}

// Handler to delete user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	err := h.userService.DeleteUser(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
