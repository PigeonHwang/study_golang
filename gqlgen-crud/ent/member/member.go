// Code generated by ent, DO NOT EDIT.

package member

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNick holds the string denoting the nick field in the database.
	FieldNick = "nick"
	// FieldTeam holds the string denoting the team field in the database.
	FieldTeam = "team"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldImg holds the string denoting the img field in the database.
	FieldImg = "img"
	// Table holds the table name of the member in the database.
	Table = "members"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNick,
	FieldTeam,
	FieldDetail,
	FieldImg,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultNick holds the default value on creation for the "nick" field.
	DefaultNick string
	// DefaultTeam holds the default value on creation for the "team" field.
	DefaultTeam string
	// DefaultDetail holds the default value on creation for the "detail" field.
	DefaultDetail string
	// DefaultImg holds the default value on creation for the "img" field.
	DefaultImg string
)
