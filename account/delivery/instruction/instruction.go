package instruction

import (
	"context"

	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func NewInstruction() instruction.Instruction {
	r := instruction.NewInstruction()
	q := query.New(context.Background())
	r = r.
		WithField(
			instruction.NewField().
				WithName("createdAt").
				WithKey("created_at").
				WithSortable(),
		).
		WithField(
			instruction.NewField().
				WithName("updatedAt").
				WithKey("updatedAt_at").
				WithSortable(),
		).
		WithField(
			instruction.NewField().
				WithName("id").
				WithKey("id").
				WithSortable().
				WithOperator(query.Equal).
				WithOperator(query.In),
		)
	r = r.WithDefaultSorts(query.NewSorts(q).
		WithSort(query.NewSort(q).
			SortBy("updated_at").
			Desc(),
		).
		WithSort(query.NewSort(q).
			SortBy("created_at").
			Desc(),
		),
	)
	return r
}
