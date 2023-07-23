package request

type ProductRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}
