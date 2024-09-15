package account

type ResetPassword struct {
	NewPassword string `json:"newPassword"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
