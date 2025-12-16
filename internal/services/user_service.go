package services

import (
	"errors"
	"spk-be/internal/database"
	"spk-be/internal/dto"
	"spk-be/internal/models"
	"spk-be/internal/repositories"
	"spk-be/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllUser() ([]models.User, error) {
	return repositories.GetAllUser()
}

func GetUserPaginate(limit int, offset int) ([]models.User, int64, error) {
	return repositories.GetPaginateUser(limit, offset)
}

func GetUserByID(id int) (*models.User, error) {
	return repositories.GetUserByID(id)
}

func GetUserName() ([]dto.UserNameDTO, error) {
	return repositories.GetUserName()
}

func CreateUser(name, username, email, password, section, role string) (*models.User, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		UUID:     uuid.NewString(),
		Name:     name,
		Username: username,
		Password: hashed,
		Email:    email,
		Section:  section,
		Role:     role,
	}

	err = repositories.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(id int, data map[string]interface{}) (*models.User, error) {
	return repositories.UpdateUser(id, data)
}

func DeleteUser(id int) (*models.User, error) {
	return repositories.DeleteUser(id)
}

func IsUsernameExsist(username string) bool {
	var user models.User
	err := database.DB.Where("username= ?", username).First(&user).Error

	if err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}

	return true
}
