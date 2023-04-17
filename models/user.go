package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
