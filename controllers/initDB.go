package controllers

import "gorm.io/gorm"

type Controller struct {
	Masterdb *gorm.DB
}

func New(db *gorm.DB) *Controller {
	return &Controller{
		Masterdb: db,
	}
}
