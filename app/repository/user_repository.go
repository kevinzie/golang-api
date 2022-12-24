package repository

import (
	"project/scratch/app/models"
	"project/scratch/pkg/config"
)

func GetUsers() []models.Users {
	var users []models.Users

	//config.Database.First(&users)
	config.Database.Preload("Profiles").Find(&users)
	return users
}

func CreateUser() {
	var users models.Users
	config.Database.Model(users).Create(map[string]interface{}{
		"FirstName": data["FirstName"],
	})
}
