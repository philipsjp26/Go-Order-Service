package entity

type (
	UserServiceResponseEntity struct {
		Code    int16  `json:"code"`
		Message string `json:"message"`
		Data    bool   `json:"data"`
	}
	UserServiceRequestEntity struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
