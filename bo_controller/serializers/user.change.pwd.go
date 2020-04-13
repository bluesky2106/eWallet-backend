package serializers

// UserChangePwdReq : struct
type UserChangePwdReq struct {
	OldPassword        string `json:"OldPassword"`
	NewPassword        string `json:"NewPassword"`
	ConfirmNewPassword string `json:"ConfirmNewPassword"`
}
