package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	RedirectTo   string
	RandomString string
}
