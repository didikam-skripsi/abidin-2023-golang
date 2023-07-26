package request

type AttributeRequest struct {
	Name       string `validate:"required" json:"name"`
	Value      string `validate:"" json:"value"`
	RangeStart int    `validate:"" json:"range_start"`
	RangeEnd   int    `validate:"" json:"range_end"`
}
