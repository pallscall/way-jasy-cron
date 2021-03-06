// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"way-jasy-cron/cron-logger/internal/model/ent/logger"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// LoggerCreate is the builder for creating a Logger entity.
type LoggerCreate struct {
	config
	mutation *LoggerMutation
	hooks    []Hook
}

// SetLog sets the "log" field.
func (lc *LoggerCreate) SetLog(s string) *LoggerCreate {
	lc.mutation.SetLog(s)
	return lc
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (lc *LoggerCreate) SetNillableLog(s *string) *LoggerCreate {
	if s != nil {
		lc.SetLog(*s)
	}
	return lc
}

// SetOperator sets the "operator" field.
func (lc *LoggerCreate) SetOperator(s string) *LoggerCreate {
	lc.mutation.SetOperator(s)
	return lc
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (lc *LoggerCreate) SetNillableOperator(s *string) *LoggerCreate {
	if s != nil {
		lc.SetOperator(*s)
	}
	return lc
}

// SetCtime sets the "ctime" field.
func (lc *LoggerCreate) SetCtime(t time.Time) *LoggerCreate {
	lc.mutation.SetCtime(t)
	return lc
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (lc *LoggerCreate) SetNillableCtime(t *time.Time) *LoggerCreate {
	if t != nil {
		lc.SetCtime(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LoggerCreate) SetID(i int) *LoggerCreate {
	lc.mutation.SetID(i)
	return lc
}

// Mutation returns the LoggerMutation object of the builder.
func (lc *LoggerCreate) Mutation() *LoggerMutation {
	return lc.mutation
}

// Save creates the Logger in the database.
func (lc *LoggerCreate) Save(ctx context.Context) (*Logger, error) {
	var (
		err  error
		node *Logger
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoggerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			node, err = lc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LoggerCreate) SaveX(ctx context.Context) *Logger {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (lc *LoggerCreate) defaults() {
	if _, ok := lc.mutation.Log(); !ok {
		v := logger.DefaultLog
		lc.mutation.SetLog(v)
	}
	if _, ok := lc.mutation.Operator(); !ok {
		v := logger.DefaultOperator
		lc.mutation.SetOperator(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LoggerCreate) check() error {
	if _, ok := lc.mutation.Log(); !ok {
		return &ValidationError{Name: "log", err: errors.New("ent: missing required field \"log\"")}
	}
	if _, ok := lc.mutation.Operator(); !ok {
		return &ValidationError{Name: "operator", err: errors.New("ent: missing required field \"operator\"")}
	}
	if v, ok := lc.mutation.ID(); ok {
		if err := logger.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (lc *LoggerCreate) sqlSave(ctx context.Context) (*Logger, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (lc *LoggerCreate) createSpec() (*Logger, *sqlgraph.CreateSpec) {
	var (
		_node = &Logger{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: logger.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logger.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.Log(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldLog,
		})
		_node.Log = value
	}
	if value, ok := lc.mutation.Operator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logger.FieldOperator,
		})
		_node.Operator = value
	}
	if value, ok := lc.mutation.Ctime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logger.FieldCtime,
		})
		_node.Ctime = value
	}
	return _node, _spec
}

// LoggerCreateBulk is the builder for creating many Logger entities in bulk.
type LoggerCreateBulk struct {
	config
	builders []*LoggerCreate
}

// Save creates the Logger entities in the database.
func (lcb *LoggerCreateBulk) Save(ctx context.Context) ([]*Logger, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Logger, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoggerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LoggerCreateBulk) SaveX(ctx context.Context) []*Logger {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
