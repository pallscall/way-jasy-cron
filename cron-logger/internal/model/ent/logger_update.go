// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"
	"way-jasy-cron/cron-logger/internal/model/ent/logger"
	"way-jasy-cron/cron-logger/internal/model/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// LoggerUpdate is the builder for updating Logger entities.
type LoggerUpdate struct {
	config
	hooks    []Hook
	mutation *LoggerMutation
}

// Where adds a new predicate for the LoggerUpdate builder.
func (lu *LoggerUpdate) Where(ps ...predicate.Logger) *LoggerUpdate {
	lu.mutation.predicates = append(lu.mutation.predicates, ps...)
	return lu
}

// SetLog sets the "log" field.
func (lu *LoggerUpdate) SetLog(s string) *LoggerUpdate {
	lu.mutation.SetLog(s)
	return lu
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (lu *LoggerUpdate) SetNillableLog(s *string) *LoggerUpdate {
	if s != nil {
		lu.SetLog(*s)
	}
	return lu
}

// SetOperator sets the "operator" field.
func (lu *LoggerUpdate) SetOperator(s string) *LoggerUpdate {
	lu.mutation.SetOperator(s)
	return lu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (lu *LoggerUpdate) SetNillableOperator(s *string) *LoggerUpdate {
	if s != nil {
		lu.SetOperator(*s)
	}
	return lu
}

// SetCtime sets the "ctime" field.
func (lu *LoggerUpdate) SetCtime(t time.Time) *LoggerUpdate {
	lu.mutation.SetCtime(t)
	return lu
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (lu *LoggerUpdate) SetNillableCtime(t *time.Time) *LoggerUpdate {
	if t != nil {
		lu.SetCtime(*t)
	}
	return lu
}

// ClearCtime clears the value of the "ctime" field.
func (lu *LoggerUpdate) ClearCtime() *LoggerUpdate {
	lu.mutation.ClearCtime()
	return lu
}

// Mutation returns the LoggerMutation object of the builder.
func (lu *LoggerUpdate) Mutation() *LoggerMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LoggerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoggerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LoggerUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LoggerUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LoggerUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LoggerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logger.Table,
			Columns: logger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logger.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldLog,
		})
	}
	if value, ok := lu.mutation.Operator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldOperator,
		})
	}
	if value, ok := lu.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logger.FieldCtime,
		})
	}
	if lu.mutation.CtimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: logger.FieldCtime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logger.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LoggerUpdateOne is the builder for updating a single Logger entity.
type LoggerUpdateOne struct {
	config
	hooks    []Hook
	mutation *LoggerMutation
}

// SetLog sets the "log" field.
func (luo *LoggerUpdateOne) SetLog(s string) *LoggerUpdateOne {
	luo.mutation.SetLog(s)
	return luo
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (luo *LoggerUpdateOne) SetNillableLog(s *string) *LoggerUpdateOne {
	if s != nil {
		luo.SetLog(*s)
	}
	return luo
}

// SetOperator sets the "operator" field.
func (luo *LoggerUpdateOne) SetOperator(s string) *LoggerUpdateOne {
	luo.mutation.SetOperator(s)
	return luo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (luo *LoggerUpdateOne) SetNillableOperator(s *string) *LoggerUpdateOne {
	if s != nil {
		luo.SetOperator(*s)
	}
	return luo
}

// SetCtime sets the "ctime" field.
func (luo *LoggerUpdateOne) SetCtime(t time.Time) *LoggerUpdateOne {
	luo.mutation.SetCtime(t)
	return luo
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (luo *LoggerUpdateOne) SetNillableCtime(t *time.Time) *LoggerUpdateOne {
	if t != nil {
		luo.SetCtime(*t)
	}
	return luo
}

// ClearCtime clears the value of the "ctime" field.
func (luo *LoggerUpdateOne) ClearCtime() *LoggerUpdateOne {
	luo.mutation.ClearCtime()
	return luo
}

// Mutation returns the LoggerMutation object of the builder.
func (luo *LoggerUpdateOne) Mutation() *LoggerMutation {
	return luo.mutation
}

// Save executes the query and returns the updated Logger entity.
func (luo *LoggerUpdateOne) Save(ctx context.Context) (*Logger, error) {
	var (
		err  error
		node *Logger
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoggerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LoggerUpdateOne) SaveX(ctx context.Context) *Logger {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LoggerUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LoggerUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LoggerUpdateOne) sqlSave(ctx context.Context) (_node *Logger, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logger.Table,
			Columns: logger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logger.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Logger.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := luo.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldLog,
		})
	}
	if value, ok := luo.mutation.Operator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldOperator,
		})
	}
	if value, ok := luo.mutation.Ctime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logger.FieldCtime,
		})
	}
	if luo.mutation.CtimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: logger.FieldCtime,
		})
	}
	_node = &Logger{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logger.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}