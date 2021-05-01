package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Machine holds the schema definition for the Machine entity.
type Machine struct {
	ent.Schema
}

// Fields of the Machine.
func (Machine) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("host").Default(""),
		field.Int("port").Default(0),
		field.String("username").Default(""),
		field.String("password").Default(""),
		field.String("comment").Default(""),
		field.String("command").Default(""),
		field.String("creator").Default(""),
		field.Int("status").Default(0),
		field.Time("ctime").Optional(),
		field.Time("mtime").Optional(),
	}
}

// Edges of the Machine.
func (Machine) Edges() []ent.Edge {
	return nil
}
