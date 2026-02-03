package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// Function to generate unique ID with prefix
func GenerateID(prefix string, db *gorm.DB, table interface{}, column string) (string, error) {
	
	lastRecord := map[string]interface{}{}

	err := db.Model(table).Select(column).Order(column + " DESC").Limit(1).Find(&lastRecord).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	var lastID string
	if val, ok := lastRecord[column]; ok && val != nil {
		lastID = val.(string)
	}

	if lastID == "" {
		return fmt.Sprintf("%s001", prefix), nil
	}

	numStr := strings.TrimPrefix(lastID, prefix)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%03d", prefix, num+1), nil
}
