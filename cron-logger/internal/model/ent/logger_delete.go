// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"way-jasy-cron/cron-logger/internal/model/ent/logger"
	"way-jasy-cron/cron-logger/internal/model/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// LoggerDelete is the builder for deleting a Logger entity.
type LoggerDelete struct {
	config
	hooks    []Hook
	mutation *LoggerMutation
}

// Where adds a new predicate to the LoggerDelete builder.
func (ld *LoggerDelete) Where(ps ...predicate.Logger) *LoggerDelete {
	ld.mutation.predicates = append(ld.mutation.predicates, ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LoggerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ld.hooks) == 0 {
		affected, err = ld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoggerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ld.mutation = mutation
			affected, err = ld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ld.hooks) - 1; i >= 0; i-- {
			mut = ld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LoggerDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LoggerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: logger.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logger.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
}

// LoggerDeleteOne is the builder for deleting a single Logger entity.
type LoggerDeleteOne struct {
	ld *LoggerDelete
}

// Exec executes the deletion query.
func (ldo *LoggerDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logger.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LoggerDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}
