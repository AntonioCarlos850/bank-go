package dto

type SuccessTransfer struct {
	ToBalance   float64 `json:"to_balance"`
	FromBalance float64 `json:"from_balance"`
}
