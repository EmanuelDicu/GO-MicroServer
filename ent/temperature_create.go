// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tema-sprc-go/ent/temperature"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemperatureCreate is the builder for creating a Temperature entity.
type TemperatureCreate struct {
	config
	mutation *TemperatureMutation
	hooks    []Hook
}

// SetIDOras sets the "id_oras" field.
func (tc *TemperatureCreate) SetIDOras(i int) *TemperatureCreate {
	tc.mutation.SetIDOras(i)
	return tc
}

// SetValoare sets the "Valoare" field.
func (tc *TemperatureCreate) SetValoare(f float64) *TemperatureCreate {
	tc.mutation.SetValoare(f)
	return tc
}

// SetTimestamp sets the "Timestamp" field.
func (tc *TemperatureCreate) SetTimestamp(t time.Time) *TemperatureCreate {
	tc.mutation.SetTimestamp(t)
	return tc
}

// Mutation returns the TemperatureMutation object of the builder.
func (tc *TemperatureCreate) Mutation() *TemperatureMutation {
	return tc.mutation
}

// Save creates the Temperature in the database.
func (tc *TemperatureCreate) Save(ctx context.Context) (*Temperature, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemperatureCreate) SaveX(ctx context.Context) *Temperature {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TemperatureCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TemperatureCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TemperatureCreate) check() error {
	if _, ok := tc.mutation.IDOras(); !ok {
		return &ValidationError{Name: "id_oras", err: errors.New(`ent: missing required field "Temperature.id_oras"`)}
	}
	if _, ok := tc.mutation.Valoare(); !ok {
		return &ValidationError{Name: "Valoare", err: errors.New(`ent: missing required field "Temperature.Valoare"`)}
	}
	if _, ok := tc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "Temperature.Timestamp"`)}
	}
	return nil
}

func (tc *TemperatureCreate) sqlSave(ctx context.Context) (*Temperature, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TemperatureCreate) createSpec() (*Temperature, *sqlgraph.CreateSpec) {
	var (
		_node = &Temperature{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(temperature.Table, sqlgraph.NewFieldSpec(temperature.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.IDOras(); ok {
		_spec.SetField(temperature.FieldIDOras, field.TypeInt, value)
		_node.IDOras = value
	}
	if value, ok := tc.mutation.Valoare(); ok {
		_spec.SetField(temperature.FieldValoare, field.TypeFloat64, value)
		_node.Valoare = value
	}
	if value, ok := tc.mutation.Timestamp(); ok {
		_spec.SetField(temperature.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	return _node, _spec
}

// TemperatureCreateBulk is the builder for creating many Temperature entities in bulk.
type TemperatureCreateBulk struct {
	config
	err      error
	builders []*TemperatureCreate
}

// Save creates the Temperature entities in the database.
func (tcb *TemperatureCreateBulk) Save(ctx context.Context) ([]*Temperature, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Temperature, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemperatureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TemperatureCreateBulk) SaveX(ctx context.Context) []*Temperature {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TemperatureCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TemperatureCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
