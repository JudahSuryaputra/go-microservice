package dto

type GetUserByIDResponse struct {
	FullName    string
	PhoneNumber string
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateUserResponse struct {
}
