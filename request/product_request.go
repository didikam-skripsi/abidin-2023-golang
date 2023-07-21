package request

type ProductRequest struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}
