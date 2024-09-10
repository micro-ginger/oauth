package instruction

import (
	"context"

	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

var integrated instruction.Instruction

func NewInstructionIntegrated() instruction.Instruction {
	if integrated != nil {
		return integrated
	}
	integrated = instruction.NewInstruction()
	q := query.New(context.Background())
	integrated = integrated.
		WithField(
			instruction.NewField().
				WithName("createdAt").
				WithKey("acc.created_at").
				WithSortable(),
		).
		WithField(
			instruction.NewField().
				WithName("updatedAt").
				WithKey("p.updatedAt_at").
				WithSortable(),
		).
		WithField(
			instruction.NewField().
				WithName("id").
				WithKey("p.id").
				WithSortable().
				WithOperator(query.Equal).
				WithOperator(query.In),
		)
	integrated = integrated.WithDefaultSorts(query.NewSorts(q).
		WithSort(query.NewSort(q).
			SortBy("p.updated_at").
			Desc(),
		).
		WithSort(query.NewSort(q).
			SortBy("acc.created_at").
			Desc(),
		),
	)
	return integrated
}
