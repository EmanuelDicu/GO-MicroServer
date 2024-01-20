// Code generated by ent, DO NOT EDIT.

package temperature

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the temperature type in the database.
	Label = "temperature"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIDOras holds the string denoting the id_oras field in the database.
	FieldIDOras = "id_oras"
	// FieldValoare holds the string denoting the valoare field in the database.
	FieldValoare = "valoare"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// Table holds the table name of the temperature in the database.
	Table = "temperatures"
)

// Columns holds all SQL columns for temperature fields.
var Columns = []string{
	FieldID,
	FieldIDOras,
	FieldValoare,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "temperatures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"id_oras",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Temperature queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIDOras orders the results by the id_oras field.
func ByIDOras(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIDOras, opts...).ToFunc()
}

// ByValoare orders the results by the Valoare field.
func ByValoare(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValoare, opts...).ToFunc()
}

// ByTimestamp orders the results by the Timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}