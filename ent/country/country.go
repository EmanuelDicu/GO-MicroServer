// Code generated by ent, DO NOT EDIT.

package country

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the country type in the database.
	Label = "country"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumeTara holds the string denoting the nume_tara field in the database.
	FieldNumeTara = "nume_tara"
	// FieldLatitudine holds the string denoting the latitudine field in the database.
	FieldLatitudine = "latitudine"
	// FieldLongitudine holds the string denoting the longitudine field in the database.
	FieldLongitudine = "longitudine"
	// EdgeCities holds the string denoting the cities edge name in mutations.
	EdgeCities = "Cities"
	// Table holds the table name of the country in the database.
	Table = "countries"
	// CitiesTable is the table that holds the Cities relation/edge.
	CitiesTable = "cities"
	// CitiesInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CitiesInverseTable = "cities"
	// CitiesColumn is the table column denoting the Cities relation/edge.
	CitiesColumn = "id_tara"
)

// Columns holds all SQL columns for country fields.
var Columns = []string{
	FieldID,
	FieldNumeTara,
	FieldLatitudine,
	FieldLongitudine,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Country queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumeTara orders the results by the Nume_tara field.
func ByNumeTara(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumeTara, opts...).ToFunc()
}

// ByLatitudine orders the results by the Latitudine field.
func ByLatitudine(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitudine, opts...).ToFunc()
}

// ByLongitudine orders the results by the Longitudine field.
func ByLongitudine(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitudine, opts...).ToFunc()
}

// ByCitiesCount orders the results by Cities count.
func ByCitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCitiesStep(), opts...)
	}
}

// ByCities orders the results by Cities terms.
func ByCities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CitiesTable, CitiesColumn),
	)
}
