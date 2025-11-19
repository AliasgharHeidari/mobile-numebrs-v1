package model

type MobileNumber struct {
	Number string `json:"number"`
	Type string `json:"type"`
	IsActive bool `json:"isActive"`
	CountryCode string `json:"countryCode"`
}

type AddMobileNumberSuccessResponse struct {
	Message string `json:"message" example:"Mobile number added successfully"`
}

type DeleteMobileNumberSuccessResponse struct {
	Message string `json:"message" example:"Mobile number deleted successfully"`
}