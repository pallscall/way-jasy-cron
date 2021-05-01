package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("name").Default(""),
		field.String("creator").Default(""),
		field.String("url").Default(""),
		field.String("spec").Default(""),
		field.String("comment").Default(""),
		field.String("updater").Default(""),
		field.String("method").Default("GET"),
		field.String("body").Default(""),
		field.String("header").Default(""),
		field.Int("count").Default(0),
		field.Int("retry").Default(0),
		field.Int("retry_temp").Default(0),
		field.Int("status").Default(0),
		field.Time("ctime").Optional(),
		field.Time("mtime").Optional(),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
