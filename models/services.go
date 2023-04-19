package models

import "gorm.io/gorm"

type Services struct {
	DbPsgl *gorm.DB
}
