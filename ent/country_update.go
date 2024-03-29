// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"tema-sprc-go/ent/city"
	"tema-sprc-go/ent/country"
	"tema-sprc-go/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CountryUpdate is the builder for updating Country entities.
type CountryUpdate struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryUpdate builder.
func (cu *CountryUpdate) Where(ps ...predicate.Country) *CountryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetNumeTara sets the "Nume_tara" field.
func (cu *CountryUpdate) SetNumeTara(s string) *CountryUpdate {
	cu.mutation.SetNumeTara(s)
	return cu
}

// SetNillableNumeTara sets the "Nume_tara" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableNumeTara(s *string) *CountryUpdate {
	if s != nil {
		cu.SetNumeTara(*s)
	}
	return cu
}

// SetLatitudine sets the "Latitudine" field.
func (cu *CountryUpdate) SetLatitudine(f float64) *CountryUpdate {
	cu.mutation.ResetLatitudine()
	cu.mutation.SetLatitudine(f)
	return cu
}

// SetNillableLatitudine sets the "Latitudine" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableLatitudine(f *float64) *CountryUpdate {
	if f != nil {
		cu.SetLatitudine(*f)
	}
	return cu
}

// AddLatitudine adds f to the "Latitudine" field.
func (cu *CountryUpdate) AddLatitudine(f float64) *CountryUpdate {
	cu.mutation.AddLatitudine(f)
	return cu
}

// SetLongitudine sets the "Longitudine" field.
func (cu *CountryUpdate) SetLongitudine(f float64) *CountryUpdate {
	cu.mutation.ResetLongitudine()
	cu.mutation.SetLongitudine(f)
	return cu
}

// SetNillableLongitudine sets the "Longitudine" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableLongitudine(f *float64) *CountryUpdate {
	if f != nil {
		cu.SetLongitudine(*f)
	}
	return cu
}

// AddLongitudine adds f to the "Longitudine" field.
func (cu *CountryUpdate) AddLongitudine(f float64) *CountryUpdate {
	cu.mutation.AddLongitudine(f)
	return cu
}

// AddCityIDs adds the "Cities" edge to the City entity by IDs.
func (cu *CountryUpdate) AddCityIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddCityIDs(ids...)
	return cu
}

// AddCities adds the "Cities" edges to the City entity.
func (cu *CountryUpdate) AddCities(c ...*City) *CountryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCityIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cu *CountryUpdate) Mutation() *CountryMutation {
	return cu.mutation
}

// ClearCities clears all "Cities" edges to the City entity.
func (cu *CountryUpdate) ClearCities() *CountryUpdate {
	cu.mutation.ClearCities()
	return cu
}

// RemoveCityIDs removes the "Cities" edge to City entities by IDs.
func (cu *CountryUpdate) RemoveCityIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveCityIDs(ids...)
	return cu
}

// RemoveCities removes "Cities" edges to City entities.
func (cu *CountryUpdate) RemoveCities(c ...*City) *CountryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCityIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.NumeTara(); ok {
		_spec.SetField(country.FieldNumeTara, field.TypeString, value)
	}
	if value, ok := cu.mutation.Latitudine(); ok {
		_spec.SetField(country.FieldLatitudine, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLatitudine(); ok {
		_spec.AddField(country.FieldLatitudine, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.Longitudine(); ok {
		_spec.SetField(country.FieldLongitudine, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLongitudine(); ok {
		_spec.AddField(country.FieldLongitudine, field.TypeFloat64, value)
	}
	if cu.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !cu.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CountryUpdateOne is the builder for updating a single Country entity.
type CountryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryMutation
}

// SetNumeTara sets the "Nume_tara" field.
func (cuo *CountryUpdateOne) SetNumeTara(s string) *CountryUpdateOne {
	cuo.mutation.SetNumeTara(s)
	return cuo
}

// SetNillableNumeTara sets the "Nume_tara" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableNumeTara(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetNumeTara(*s)
	}
	return cuo
}

// SetLatitudine sets the "Latitudine" field.
func (cuo *CountryUpdateOne) SetLatitudine(f float64) *CountryUpdateOne {
	cuo.mutation.ResetLatitudine()
	cuo.mutation.SetLatitudine(f)
	return cuo
}

// SetNillableLatitudine sets the "Latitudine" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableLatitudine(f *float64) *CountryUpdateOne {
	if f != nil {
		cuo.SetLatitudine(*f)
	}
	return cuo
}

// AddLatitudine adds f to the "Latitudine" field.
func (cuo *CountryUpdateOne) AddLatitudine(f float64) *CountryUpdateOne {
	cuo.mutation.AddLatitudine(f)
	return cuo
}

// SetLongitudine sets the "Longitudine" field.
func (cuo *CountryUpdateOne) SetLongitudine(f float64) *CountryUpdateOne {
	cuo.mutation.ResetLongitudine()
	cuo.mutation.SetLongitudine(f)
	return cuo
}

// SetNillableLongitudine sets the "Longitudine" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableLongitudine(f *float64) *CountryUpdateOne {
	if f != nil {
		cuo.SetLongitudine(*f)
	}
	return cuo
}

// AddLongitudine adds f to the "Longitudine" field.
func (cuo *CountryUpdateOne) AddLongitudine(f float64) *CountryUpdateOne {
	cuo.mutation.AddLongitudine(f)
	return cuo
}

// AddCityIDs adds the "Cities" edge to the City entity by IDs.
func (cuo *CountryUpdateOne) AddCityIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddCityIDs(ids...)
	return cuo
}

// AddCities adds the "Cities" edges to the City entity.
func (cuo *CountryUpdateOne) AddCities(c ...*City) *CountryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCityIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cuo *CountryUpdateOne) Mutation() *CountryMutation {
	return cuo.mutation
}

// ClearCities clears all "Cities" edges to the City entity.
func (cuo *CountryUpdateOne) ClearCities() *CountryUpdateOne {
	cuo.mutation.ClearCities()
	return cuo
}

// RemoveCityIDs removes the "Cities" edge to City entities by IDs.
func (cuo *CountryUpdateOne) RemoveCityIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveCityIDs(ids...)
	return cuo
}

// RemoveCities removes "Cities" edges to City entities.
func (cuo *CountryUpdateOne) RemoveCities(c ...*City) *CountryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCityIDs(ids...)
}

// Where appends a list predicates to the CountryUpdate builder.
func (cuo *CountryUpdateOne) Where(ps ...predicate.Country) *CountryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountryUpdateOne) Select(field string, fields ...string) *CountryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Country entity.
func (cuo *CountryUpdateOne) Save(ctx context.Context) (*Country, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountryUpdateOne) SaveX(ctx context.Context) *Country {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CountryUpdateOne) sqlSave(ctx context.Context) (_node *Country, err error) {
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for _, f := range fields {
			if !country.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.NumeTara(); ok {
		_spec.SetField(country.FieldNumeTara, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Latitudine(); ok {
		_spec.SetField(country.FieldLatitudine, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLatitudine(); ok {
		_spec.AddField(country.FieldLatitudine, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.Longitudine(); ok {
		_spec.SetField(country.FieldLongitudine, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLongitudine(); ok {
		_spec.AddField(country.FieldLongitudine, field.TypeFloat64, value)
	}
	if cuo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCitiesIDs(); len(nodes) > 0 && !cuo.mutation.CitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.CitiesTable,
			Columns: []string{country.CitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
