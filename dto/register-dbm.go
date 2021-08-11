package dto

// client post
type RegisterDTO struct{
	Name string `json:"name" form:"name" validate:"min:1"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:8"`
}