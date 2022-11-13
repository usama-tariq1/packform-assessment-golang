package controllers

import (
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

var ()

type Controller struct {
}

func Lamp() *gorm.DB {
	return models.DB
}
