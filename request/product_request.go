package request

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ProductRequestID struct {
	ID uint `uri:"id" binding:"required"`
}
