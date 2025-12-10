package seeder

import (
	"log"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/database"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
)

func SeedUser() {

	DB := database.GetDB()

	user := model.User{
		Name:          "Ethan",
		FamilyName:    "Winters",
		Age:           27,
		IsMarried:     true,
		MobileNumbers:[]model.MobileNumber{
		{
			Number: "09335433443",
			CountryCode: "98",
			IsActive: false,
			Type: "debit",
		},
		},
	}

	err := DB.Create(&user).Error
	if err != nil {
		log.Fatal("failed to seed user to database")
	}

}
