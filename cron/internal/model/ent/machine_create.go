// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"way-jasy-cron/cron/internal/model/ent/machine"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MachineCreate is the builder for creating a Machine entity.
type MachineCreate struct {
	config
	mutation *MachineMutation
	hooks    []Hook
}

// SetHost sets the "host" field.
func (mc *MachineCreate) SetHost(s string) *MachineCreate {
	mc.mutation.SetHost(s)
	return mc
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mc *MachineCreate) SetNillableHost(s *string) *MachineCreate {
	if s != nil {
		mc.SetHost(*s)
	}
	return mc
}

// SetPort sets the "port" field.
func (mc *MachineCreate) SetPort(i int) *MachineCreate {
	mc.mutation.SetPort(i)
	return mc
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (mc *MachineCreate) SetNillablePort(i *int) *MachineCreate {
	if i != nil {
		mc.SetPort(*i)
	}
	return mc
}

// SetUsername sets the "username" field.
func (mc *MachineCreate) SetUsername(s string) *MachineCreate {
	mc.mutation.SetUsername(s)
	return mc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (mc *MachineCreate) SetNillableUsername(s *string) *MachineCreate {
	if s != nil {
		mc.SetUsername(*s)
	}
	return mc
}

// SetPassword sets the "password" field.
func (mc *MachineCreate) SetPassword(s string) *MachineCreate {
	mc.mutation.SetPassword(s)
	return mc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mc *MachineCreate) SetNillablePassword(s *string) *MachineCreate {
	if s != nil {
		mc.SetPassword(*s)
	}
	return mc
}

// SetComment sets the "comment" field.
func (mc *MachineCreate) SetComment(s string) *MachineCreate {
	mc.mutation.SetComment(s)
	return mc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (mc *MachineCreate) SetNillableComment(s *string) *MachineCreate {
	if s != nil {
		mc.SetComment(*s)
	}
	return mc
}

// SetCommand sets the "command" field.
func (mc *MachineCreate) SetCommand(s string) *MachineCreate {
	mc.mutation.SetCommand(s)
	return mc
}

// SetNillableCommand sets the "command" field if the given value is not nil.
func (mc *MachineCreate) SetNillableCommand(s *string) *MachineCreate {
	if s != nil {
		mc.SetCommand(*s)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MachineCreate) SetStatus(i int) *MachineCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MachineCreate) SetNillableStatus(i *int) *MachineCreate {
	if i != nil {
		mc.SetStatus(*i)
	}
	return mc
}

// SetCtime sets the "ctime" field.
func (mc *MachineCreate) SetCtime(t time.Time) *MachineCreate {
	mc.mutation.SetCtime(t)
	return mc
}

// SetNillableCtime sets the "ctime" field if the given value is not nil.
func (mc *MachineCreate) SetNillableCtime(t *time.Time) *MachineCreate {
	if t != nil {
		mc.SetCtime(*t)
	}
	return mc
}

// SetMtime sets the "mtime" field.
func (mc *MachineCreate) SetMtime(t time.Time) *MachineCreate {
	mc.mutation.SetMtime(t)
	return mc
}

// SetNillableMtime sets the "mtime" field if the given value is not nil.
func (mc *MachineCreate) SetNillableMtime(t *time.Time) *MachineCreate {
	if t != nil {
		mc.SetMtime(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MachineCreate) SetID(i int) *MachineCreate {
	mc.mutation.SetID(i)
	return mc
}

// Mutation returns the MachineMutation object of the builder.
func (mc *MachineCreate) Mutation() *MachineMutation {
	return mc.mutation
}

// Save creates the Machine in the database.
func (mc *MachineCreate) Save(ctx context.Context) (*Machine, error) {
	var (
		err  error
		node *Machine
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MachineCreate) SaveX(ctx context.Context) *Machine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mc *MachineCreate) defaults() {
	if _, ok := mc.mutation.Host(); !ok {
		v := machine.DefaultHost
		mc.mutation.SetHost(v)
	}
	if _, ok := mc.mutation.Port(); !ok {
		v := machine.DefaultPort
		mc.mutation.SetPort(v)
	}
	if _, ok := mc.mutation.Username(); !ok {
		v := machine.DefaultUsername
		mc.mutation.SetUsername(v)
	}
	if _, ok := mc.mutation.Password(); !ok {
		v := machine.DefaultPassword
		mc.mutation.SetPassword(v)
	}
	if _, ok := mc.mutation.Comment(); !ok {
		v := machine.DefaultComment
		mc.mutation.SetComment(v)
	}
	if _, ok := mc.mutation.Command(); !ok {
		v := machine.DefaultCommand
		mc.mutation.SetCommand(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := machine.DefaultStatus
		mc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MachineCreate) check() error {
	if _, ok := mc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New("ent: missing required field \"host\"")}
	}
	if _, ok := mc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New("ent: missing required field \"port\"")}
	}
	if _, ok := mc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New("ent: missing required field \"username\"")}
	}
	if _, ok := mc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if _, ok := mc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New("ent: missing required field \"comment\"")}
	}
	if _, ok := mc.mutation.Command(); !ok {
		return &ValidationError{Name: "command", err: errors.New("ent: missing required field \"command\"")}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := machine.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (mc *MachineCreate) sqlSave(ctx context.Context) (*Machine, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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

func (mc *MachineCreate) createSpec() (*Machine, *sqlgraph.CreateSpec) {
	var (
		_node = &Machine{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: machine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: machine.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Host(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldHost,
		})
		_node.Host = value
	}
	if value, ok := mc.mutation.Port(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: machine.FieldPort,
		})
		_node.Port = value
	}
	if value, ok := mc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := mc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := mc.mutation.Command(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: machine.FieldCommand,
		})
		_node.Command = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: machine.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := mc.mutation.Ctime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldCtime,
		})
		_node.Ctime = value
	}
	if value, ok := mc.mutation.Mtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: machine.FieldMtime,
		})
		_node.Mtime = value
	}
	return _node, _spec
}

// MachineCreateBulk is the builder for creating many Machine entities in bulk.
type MachineCreateBulk struct {
	config
	builders []*MachineCreate
}

// Save creates the Machine entities in the database.
func (mcb *MachineCreateBulk) Save(ctx context.Context) ([]*Machine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Machine, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MachineMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MachineCreateBulk) SaveX(ctx context.Context) []*Machine {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
