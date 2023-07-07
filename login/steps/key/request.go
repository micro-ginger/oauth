package key

type Request struct {
	Key      string `json:"key" binding:"required"`
	Password string `json:"password" binding:"required"`
}
