// Code generated by entc, DO NOT EDIT.

package machine

const (
	// Label holds the string label denoting the machine type in the database.
	Label = "machine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldCommand holds the string denoting the command field in the database.
	FieldCommand = "command"
	// FieldCreator holds the string denoting the creator field in the database.
	FieldCreator = "creator"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"

	// Table holds the table name of the machine in the database.
	Table = "machines"
)

// Columns holds all SQL columns for machine fields.
var Columns = []string{
	FieldID,
	FieldHost,
	FieldPort,
	FieldUsername,
	FieldPassword,
	FieldComment,
	FieldCommand,
	FieldCreator,
	FieldStatus,
	FieldCtime,
	FieldMtime,
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
	// DefaultHost holds the default value on creation for the "host" field.
	DefaultHost string
	// DefaultPort holds the default value on creation for the "port" field.
	DefaultPort int
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultComment holds the default value on creation for the "comment" field.
	DefaultComment string
	// DefaultCommand holds the default value on creation for the "command" field.
	DefaultCommand string
	// DefaultCreator holds the default value on creation for the "creator" field.
	DefaultCreator string
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
