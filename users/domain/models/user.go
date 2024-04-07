package users_models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
