package model

type MobileNumber struct {
	Number string `json:"number"`
	Type string `json:"type"`
	IsActive bool `json:"isActive"`
	CountryCode string `json:"countryCode"`
}
