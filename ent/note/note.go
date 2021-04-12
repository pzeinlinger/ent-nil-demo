// Code generated by entc, DO NOT EDIT.

package note

const (
	// Label holds the string label denoting the note type in the database.
	Label = "note"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldExternalReference holds the string denoting the external_reference field in the database.
	FieldExternalReference = "external_reference"
	// Table holds the table name of the note in the database.
	Table = "notes"
)

// Columns holds all SQL columns for note fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldExternalReference,
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

var (
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
)
