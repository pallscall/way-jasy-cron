// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "creator", Type: field.TypeString, Default: ""},
		{Name: "url", Type: field.TypeString, Default: ""},
		{Name: "spec", Type: field.TypeString, Default: ""},
		{Name: "comment", Type: field.TypeString, Default: ""},
		{Name: "updater", Type: field.TypeString, Default: ""},
		{Name: "method", Type: field.TypeString, Default: "GET"},
		{Name: "body", Type: field.TypeString, Default: ""},
		{Name: "header", Type: field.TypeString, Default: ""},
		{Name: "stoppable", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
		{Name: "ctime", Type: field.TypeTime, Nullable: true},
		{Name: "mtime", Type: field.TypeTime, Nullable: true},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:        "jobs",
		Columns:     JobsColumns,
		PrimaryKey:  []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JobsTable,
	}
)

func init() {
}