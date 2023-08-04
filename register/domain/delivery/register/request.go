package register

type Request[T any] struct {
	Password *string `json:"password"`

	T T `json:"extra"`
}
