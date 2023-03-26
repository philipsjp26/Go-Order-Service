package entity

type (
	OrderRequest struct {
		PrivyID  string `json:"privy_id"`
		MenuID   int16  `json:"menu_id"`
		Quantity int16  `json:"quantity"`
		Status   string `json:"status"`
	}
	OrderResponse struct {
		Id     int16  `json:"id"`
		Status string `json:"status"`
	}
)

const (
	PROCESS   = "process"
	DELIVERED = "delivered"
)

func GetPrice(menu_id int32) int32 {
	switch menu_id {
	case 1:
		return 1000
	case 2:
		return 2000
	default:
		return 0
	}
}
