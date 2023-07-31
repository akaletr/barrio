package user

import "gorm.io/gorm"

type user struct {
	conn      *gorm.DB
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	About     string `json:"about"`
	Email     string `json:"email"`
}
