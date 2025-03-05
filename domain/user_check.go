package domain

type UserCheckRequest struct {
	Email       string `form:"email" binding:"omitempty,email"`
	PhoneNumber string `form:"phone_number" binding:"omitempty"`
}

type UserCheckResponse struct {
	Exists bool `json:"exists"`
}