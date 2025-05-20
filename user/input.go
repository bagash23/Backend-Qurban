package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}