package util

import (
	"github.com/gofiber/fiber/v2"
)

func UploadFileProduct(c *fiber.Ctx, formKey string, savePath string) (string, error) {
	fileHeader, err := c.FormFile(formKey)
	if err != nil {
		return "", c.Status(400).JSON(fiber.Map{"error": "can't upload file"})
	}

	filePath := savePath + fileHeader.Filename
	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}
