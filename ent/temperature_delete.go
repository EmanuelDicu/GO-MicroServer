// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"tema-sprc-go/ent/predicate"
	"tema-sprc-go/ent/temperature"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemperatureDelete is the builder for deleting a Temperature entity.
type TemperatureDelete struct {
	config
	hooks    []Hook
	mutation *TemperatureMutation
}

// Where appends a list predicates to the TemperatureDelete builder.
func (td *TemperatureDelete) Where(ps ...predicate.Temperature) *TemperatureDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TemperatureDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TemperatureDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TemperatureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(temperature.Table, sqlgraph.NewFieldSpec(temperature.FieldID, field.TypeInt))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TemperatureDeleteOne is the builder for deleting a single Temperature entity.
type TemperatureDeleteOne struct {
	td *TemperatureDelete
}

// Where appends a list predicates to the TemperatureDelete builder.
func (tdo *TemperatureDeleteOne) Where(ps ...predicate.Temperature) *TemperatureDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TemperatureDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{temperature.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TemperatureDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}