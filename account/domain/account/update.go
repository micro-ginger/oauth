package account

import "github.com/micro-blonde/auth/account"

type UpdateStatus struct {
	AddStatus account.Status
	DelStatus account.Status
}

func (u *UpdateStatus) Add(status account.Status) {
	u.AddStatus.Add(status)
}

func (u *UpdateStatus) Del(status account.Status) {
	u.DelStatus.Add(status)
}

type UpdateInternalStatus struct {
	AddStatus account.InternalStatus
	DelStatus account.InternalStatus
}

func (u *UpdateInternalStatus) Add(status account.InternalStatus) {
	u.AddStatus.Add(status)
}

func (u *UpdateInternalStatus) Del(status account.InternalStatus) {
	u.DelStatus.Add(status)
}

type UpdatePassword struct {
	New string
}

type Update[T account.Model] struct {
	Account              *Account[T]
	UpdateInternalStatus *UpdateInternalStatus

	UpdateStatus   *UpdateStatus
	UpdatePassword *UpdatePassword
	T              T
}

func (u *Update[T]) AddStatus(status account.Status) {
	if u.UpdateStatus == nil {
		u.UpdateStatus = new(UpdateStatus)
	}
	u.UpdateStatus.Add(status)
}

func (u *Update[T]) DelStatus(status account.Status) {
	if u.UpdateStatus == nil {
		u.UpdateStatus = new(UpdateStatus)
	}
	u.UpdateStatus.Del(status)
}
