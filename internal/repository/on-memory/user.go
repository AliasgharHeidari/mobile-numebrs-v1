package onmemory

import (
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
)

var Users []model.User

func InitUsers() {

	Users = []model.User{
		{
			ID: 1,
			Name: "ali",
			FamilyName: "heidari",
			Age: 30,
			IsMarried: false,
			MobileNumbers: []model.MobileNumber{
				{
					Number: "09123456789",
					Type: "creadit",
					IsActive: true,
					CountryCode: "+98",
				},
			},
		},
	}


}