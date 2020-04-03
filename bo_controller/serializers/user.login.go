package serializers

// UserLoginReq : user login request
type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLoginResp : user login response
type UserLoginResp struct {
	Token   string `json:"Token"`
	Expired string `json:"Expired"`
}
