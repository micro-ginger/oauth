package repository

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/query"
	sqlQuery "github.com/ginger-repository/sql/query"
	"github.com/micro-ginger/oauth/account/profile/domain/profile"
	"gorm.io/gorm"
)

func (repo *repo[T]) List(q query.Query) ([]*profile.Profile[T], errors.Error) {
	q = query.NewModelsQuery(q).
		WithModelsHandlerFunc(func() any {
			return new([]*profile.Profile[T])
		})
	r, err := repo.Repository.List(q)
	if err != nil {
		return nil, err.
			WithTrace("Repository.List")
	}
	return *r.(*[]*profile.Profile[T]), nil
}

func (repo *repo[T]) Get(q query.Query) (*profile.Profile[T], errors.Error) {
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(profile.Profile[T])
		})
	r, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err.
			WithTrace("Repository.Get")
	}
	return r.(*profile.Profile[T]), nil
}

func (repo *repo[T]) GetAggregated(q query.Query) (*profile.Profile[T], errors.Error) {
	q = repo.initiateProfileFetch(q)
	repo.joinAccount(q)
	//
	q = query.NewModelQuery(q).
		WithModelHandlerFunc(func() any {
			return new(profile.Profile[T])
		})
	r, err := repo.Repository.Get(q)
	if err != nil {
		return nil, err.
			WithTrace("Repository.Get")
	}
	return r.(*profile.Profile[T]), nil
}

func (repo *repo[T]) initiateProfileFetch(q query.Query) query.Query {
	db := repo.GetDB(q).(*gorm.DB).
		Table("profiles p")
	return sqlQuery.New(q, db)
}

func (repo *repo[T]) joinAccount(q query.Query) {
	db := repo.GetDB(q).(*gorm.DB)
	db = db.Joins("inner join accounts acc on acc.id = p.id")
	q.SetDB(db)
}
