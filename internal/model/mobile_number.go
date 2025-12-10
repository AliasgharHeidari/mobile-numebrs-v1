package model

type MobileNumber struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Number      string `json:"number"`
	Type        string `json:"type"`
	IsActive    bool   `json:"isActive"`
	CountryCode string `json:"countryCode"`
	UserID      uint   `json:"userId"`
}

type AddMobileNumberSuccessResponse struct {
	Message string `json:"message" example:"Mobile number added successfully"`
}

type DeleteMobileNumberSuccessResponse struct {
	Message string `json:"message" example:"Mobile number deleted successfully"`
}
