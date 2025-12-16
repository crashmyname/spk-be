package models

type User struct {
	ID       uint   `json:"user_id" gorm:"primaryKey;column:user_id"`
	UUID     string `json:"uuid" gorm:"uniqueKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Section  string `json:"section"`
	Role     string `json:"role"`
}
