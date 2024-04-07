package accountrole

type AccountRole struct {
	AccountId    uint64
	RoleId       uint64
	IsAuthorized *bool
}

func (m *AccountRole) GetId() any {
	return nil
}
