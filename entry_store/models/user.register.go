package models

// UserRegisterReq : user register request
type UserRegisterReq struct {
	FullName        string `json:"FullName"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}
