package dto

// CreateAccountDTO is awesome
type CreateAccountDTO struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
