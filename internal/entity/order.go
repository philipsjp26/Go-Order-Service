package entity

type (
	OrderRequest struct {
		PrivyID  string `json:"privy_id"`
		MenuID   int16  `json:"menu_id"`
		Quantity int16  `json:"quantity"`
	}
	OrderResponse struct {
		Id     int16  `json:"id"`
		Status string `json:"status"`
	}
)
