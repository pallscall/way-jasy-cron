// Code generated by entc, DO NOT EDIT.

package machine

import (
	"time"
	"way-jasy-cron/cron/internal/model/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPort), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// Command applies equality check predicate on the "command" field. It's identical to CommandEQ.
func Command(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommand), v))
	})
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Ctime applies equality check predicate on the "ctime" field. It's identical to CtimeEQ.
func Ctime(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHost), v))
	})
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHost), v...))
	})
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHost), v...))
	})
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHost), v))
	})
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHost), v))
	})
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHost), v))
	})
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHost), v))
	})
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHost), v))
	})
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHost), v))
	})
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHost), v))
	})
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHost), v))
	})
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHost), v))
	})
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPort), v))
	})
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPort), v))
	})
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPort), v...))
	})
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPort), v...))
	})
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPort), v))
	})
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPort), v))
	})
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPort), v))
	})
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPort), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComment), v))
	})
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldComment), v...))
	})
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldComment), v...))
	})
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComment), v))
	})
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComment), v))
	})
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComment), v))
	})
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComment), v))
	})
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComment), v))
	})
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComment), v))
	})
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComment), v))
	})
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComment), v))
	})
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComment), v))
	})
}

// CommandEQ applies the EQ predicate on the "command" field.
func CommandEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommand), v))
	})
}

// CommandNEQ applies the NEQ predicate on the "command" field.
func CommandNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommand), v))
	})
}

// CommandIn applies the In predicate on the "command" field.
func CommandIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommand), v...))
	})
}

// CommandNotIn applies the NotIn predicate on the "command" field.
func CommandNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommand), v...))
	})
}

// CommandGT applies the GT predicate on the "command" field.
func CommandGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommand), v))
	})
}

// CommandGTE applies the GTE predicate on the "command" field.
func CommandGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommand), v))
	})
}

// CommandLT applies the LT predicate on the "command" field.
func CommandLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommand), v))
	})
}

// CommandLTE applies the LTE predicate on the "command" field.
func CommandLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommand), v))
	})
}

// CommandContains applies the Contains predicate on the "command" field.
func CommandContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCommand), v))
	})
}

// CommandHasPrefix applies the HasPrefix predicate on the "command" field.
func CommandHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCommand), v))
	})
}

// CommandHasSuffix applies the HasSuffix predicate on the "command" field.
func CommandHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCommand), v))
	})
}

// CommandEqualFold applies the EqualFold predicate on the "command" field.
func CommandEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCommand), v))
	})
}

// CommandContainsFold applies the ContainsFold predicate on the "command" field.
func CommandContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCommand), v))
	})
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreator), v))
	})
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreator), v...))
	})
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreator), v...))
	})
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreator), v))
	})
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreator), v))
	})
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreator), v))
	})
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreator), v))
	})
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreator), v))
	})
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreator), v))
	})
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreator), v))
	})
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreator), v))
	})
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreator), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// CtimeEQ applies the EQ predicate on the "ctime" field.
func CtimeEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCtime), v))
	})
}

// CtimeNEQ applies the NEQ predicate on the "ctime" field.
func CtimeNEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCtime), v))
	})
}

// CtimeIn applies the In predicate on the "ctime" field.
func CtimeIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCtime), v...))
	})
}

// CtimeNotIn applies the NotIn predicate on the "ctime" field.
func CtimeNotIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCtime), v...))
	})
}

// CtimeGT applies the GT predicate on the "ctime" field.
func CtimeGT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCtime), v))
	})
}

// CtimeGTE applies the GTE predicate on the "ctime" field.
func CtimeGTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCtime), v))
	})
}

// CtimeLT applies the LT predicate on the "ctime" field.
func CtimeLT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCtime), v))
	})
}

// CtimeLTE applies the LTE predicate on the "ctime" field.
func CtimeLTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCtime), v))
	})
}

// CtimeIsNil applies the IsNil predicate on the "ctime" field.
func CtimeIsNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCtime)))
	})
}

// CtimeNotNil applies the NotNil predicate on the "ctime" field.
func CtimeNotNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCtime)))
	})
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtime), v))
	})
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtime), v))
	})
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMtime), v...))
	})
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMtime), v...))
	})
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMtime), v))
	})
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMtime), v))
	})
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMtime), v))
	})
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMtime), v))
	})
}

// MtimeIsNil applies the IsNil predicate on the "mtime" field.
func MtimeIsNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMtime)))
	})
}

// MtimeNotNil applies the NotNil predicate on the "mtime" field.
func MtimeNotNil() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMtime)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		p(s.Not())
	})
}
