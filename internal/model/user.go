package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FamilyName string `json:"familyName"`
	Age int `json:"age"`
	IsMarried bool `json:"isMarried"`
	MobileNumbers []MobileNumber `json:"mobileNumbers"`
}