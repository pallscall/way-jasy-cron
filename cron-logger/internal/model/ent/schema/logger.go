package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Logger holds the schema definition for the Logger entity.
type Logger struct {
	ent.Schema
}

// Fields of the Logger.
func (Logger) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("log").Default(""),
		field.String("operator").Default(""),
		field.Time("ctime").Optional(),
	}
}

// Edges of the Logger.
func (Logger) Edges() []ent.Edge {
	return nil
}
