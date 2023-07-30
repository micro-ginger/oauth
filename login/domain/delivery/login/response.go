package login

type Response struct {
	Sessions map[string]*Session `json:"sessions"`
}
