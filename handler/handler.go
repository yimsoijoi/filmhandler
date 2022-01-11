package handler

import (
	"gorm.io/gorm"
)

type Handler interface {
}

type handler struct {
	pg *gorm.DB
}

func New(db *gorm.DB) Handler {
	return &handler{
		pg: db,
	}
}
