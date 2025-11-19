package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FamilyName string `json:"familyName"`
	Age int `json:"age"`
	IsMarried bool `json:"isMarried"`
	MobileNumbers []MobileNumber `json:"mobileNumbers"`
}

type StatusNotFoundResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"User not found"` 
}

type StatusBadRequestResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Invalid request body"` 
}

type StatusInternalServerErrorResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Internal server error"` 
}

type StatusUnauthorizedResponse struct {
	// Error indicates the error message
	Error string `json:"error" example:"Authentication failed"` 
}

type LoginRequest struct {
	// UserName is the username of the user
	UserName string `json:"username" example:"Aliasghar"`
	// Password is the password of the user
	Password string `json:"password" example:"1234"`
}

type LoginSuccessResponse struct {
	// Token is the JWT token
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

type CreateUserRequest struct {
		// Name is the name of the user
		Name string `json:"name" example:"John"`
		// FamilyName is the family name of the user
		FamilyName string `json:"familyName" example:"Doe"`
		// Age is the age of the user
		Age int `json:"age" example:"30"`
		// IsMarried is the marital status of the user
		IsMarried bool `json:"isMarried" example:"false"`
}

type CreateUserSuccessResponse struct {
		// Message is the success message
		Message string `json:"message" example:"User created successfully"`
		// UserID is the ID of the created user
		UserID int `json:"user_id" example:"1"`
}

type DeleteUserSuccessResponse struct {
		// Message is the success message
		Message string `json:"message" example:"User deleted successfully"`
}

type UpdateUserSuccessResponse struct {
		// Message is the success message
		Message string `json:"message" example:"User updated successfully"`
}