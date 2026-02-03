package util

import (
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx, formKey string, savePath string) (string, error) {
	fileHeader, err := c.FormFile(formKey)
	if err != nil {
	
		return "", nil
	}

	filePath := savePath + fileHeader.Filename

	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}
