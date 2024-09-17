package usecase

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/query"
)

func (uc useCase[T]) startCron() {
	uc.startMtx.RLock()
	defer uc.startMtx.RUnlock()

	ticker := time.NewTicker(time.Second * 10)

	select {
	case <-uc.closeChan:
		return
	case <-ticker.C:
		err := uc.processPendings()
		if err != nil {
			uc.logger.
				With(logger.Field{
					"error": err.Error(),
				}).
				WithTrace("repo.List").
				Errorf("error on get pending accounts")
		}
	}
}

func (uc useCase[T]) processPendings() errors.Error {
	ctx := context.Background()
	q := query.New(ctx)
	q = query.NewFilter(q).
		WithMatch(&query.Match{
			Key:      "internal_status",
			Operator: query.Greater,
			Value:    0,
		})
	q = query.NewPagination(q).
		WithSize(30)
	pendings, err := uc.repo.List(q)
	if err != nil {
		return err.WithTrace("repo.List")
	}
	for _, p := range pendings {
		if err = uc.handleInternalStatus(ctx, p); err != nil {
			return err.WithTrace("handleInternalStatus")
		}
	}
	return nil
}
