package handler

import (
	"gorm.io/gorm"
)

type ApiHandler struct {
	db *gorm.DB
}

var ()

func NewApiHandler(db *gorm.DB) *ApiHandler {
	return &ApiHandler{
		db: db,
	}
}
