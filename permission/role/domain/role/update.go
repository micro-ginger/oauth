package role

type UpdateRequest struct {
	Name string
}

func (*UpdateRequest) TableName() string {
	return "roles"
}
