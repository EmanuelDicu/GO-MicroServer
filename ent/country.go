// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"tema-sprc-go/ent/country"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Country is the model entity for the Country schema.
type Country struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NumeTara holds the value of the "Nume_tara" field.
	NumeTara string `json:"Nume_tara,omitempty"`
	// Latitudine holds the value of the "Latitudine" field.
	Latitudine float64 `json:"Latitudine,omitempty"`
	// Longitudine holds the value of the "Longitudine" field.
	Longitudine float64 `json:"Longitudine,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges        CountryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// Cities holds the value of the Cities edge.
	Cities []*City `json:"Cities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CitiesOrErr returns the Cities value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) CitiesOrErr() ([]*City, error) {
	if e.loadedTypes[0] {
		return e.Cities, nil
	}
	return nil, &NotLoadedError{edge: "Cities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldLatitudine, country.FieldLongitudine:
			values[i] = new(sql.NullFloat64)
		case country.FieldID:
			values[i] = new(sql.NullInt64)
		case country.FieldNumeTara:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case country.FieldNumeTara:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Nume_tara", values[i])
			} else if value.Valid {
				c.NumeTara = value.String
			}
		case country.FieldLatitudine:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Latitudine", values[i])
			} else if value.Valid {
				c.Latitudine = value.Float64
			}
		case country.FieldLongitudine:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Longitudine", values[i])
			} else if value.Valid {
				c.Longitudine = value.Float64
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Country.
// This includes values selected through modifiers, order, etc.
func (c *Country) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCities queries the "Cities" edge of the Country entity.
func (c *Country) QueryCities() *CityQuery {
	return NewCountryClient(c.config).QueryCities(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("Nume_tara=")
	builder.WriteString(c.NumeTara)
	builder.WriteString(", ")
	builder.WriteString("Latitudine=")
	builder.WriteString(fmt.Sprintf("%v", c.Latitudine))
	builder.WriteString(", ")
	builder.WriteString("Longitudine=")
	builder.WriteString(fmt.Sprintf("%v", c.Longitudine))
	builder.WriteByte(')')
	return builder.String()
}

// Countries is a parsable slice of Country.
type Countries []*Country
