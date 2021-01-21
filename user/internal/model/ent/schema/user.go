package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("username").Default(""),
		field.String("password").Default(""),
		field.String("tel").Default(""),
		field.String("email").Default(""),
		field.String("public_key").Default(""),
		field.String("private_key").Default(""),
		field.Time("rtime").Optional().Comment("注册时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
