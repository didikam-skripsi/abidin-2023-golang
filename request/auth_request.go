package request

type RegisterRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required"`
}
