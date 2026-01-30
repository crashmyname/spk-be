package repositories

import (
	"errors"
	"spk-be/internal/database"
	"spk-be/internal/dto"
	"spk-be/internal/models"

	"gorm.io/gorm"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error

	return users, err
}

func GetPaginateUser(limit int, offset int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := database.DB.Model(&models.User{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("name LIKE ? OR username LIKE ? OR email LIKE ?",
			like, like, like)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Order("user_id DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil

}

func GetUserName() ([]dto.UserNameDTO, error) {
	var users []dto.UserNameDTO
	err := database.DB.Model(&models.User{}).Select("name").Find(&users).Error

	return users, err
}

func GetUserByID(id int) (*models.User, error) {
	var users models.User

	if err := database.DB.Where("username = ?", id).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &users, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(id int, data map[string]interface{}) (*models.User, error) {
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&user).Updates(data).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func DeleteUser(id int) (*models.User, error) {
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func ExportUser() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
