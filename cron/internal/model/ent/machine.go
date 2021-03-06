// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"way-jasy-cron/cron/internal/model/ent/machine"

	"github.com/facebook/ent/dialect/sql"
)

// Machine is the model entity for the Machine schema.
type Machine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Command holds the value of the "command" field.
	Command string `json:"command,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Ctime holds the value of the "ctime" field.
	Ctime time.Time `json:"ctime,omitempty"`
	// Mtime holds the value of the "mtime" field.
	Mtime time.Time `json:"mtime,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Machine) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case machine.FieldID, machine.FieldPort, machine.FieldStatus:
			values[i] = &sql.NullInt64{}
		case machine.FieldHost, machine.FieldUsername, machine.FieldPassword, machine.FieldComment, machine.FieldCommand, machine.FieldCreator:
			values[i] = &sql.NullString{}
		case machine.FieldCtime, machine.FieldMtime:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Machine", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Machine fields.
func (m *Machine) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case machine.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case machine.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				m.Host = value.String
			}
		case machine.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				m.Port = int(value.Int64)
			}
		case machine.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				m.Username = value.String
			}
		case machine.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				m.Password = value.String
			}
		case machine.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				m.Comment = value.String
			}
		case machine.FieldCommand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field command", values[i])
			} else if value.Valid {
				m.Command = value.String
			}
		case machine.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				m.Creator = value.String
			}
		case machine.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = int(value.Int64)
			}
		case machine.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				m.Ctime = value.Time
			}
		case machine.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				m.Mtime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Machine.
// Note that you need to call Machine.Unwrap() before calling this method if this Machine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Machine) Update() *MachineUpdateOne {
	return (&MachineClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Machine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Machine) Unwrap() *Machine {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Machine is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Machine) String() string {
	var builder strings.Builder
	builder.WriteString("Machine(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", host=")
	builder.WriteString(m.Host)
	builder.WriteString(", port=")
	builder.WriteString(fmt.Sprintf("%v", m.Port))
	builder.WriteString(", username=")
	builder.WriteString(m.Username)
	builder.WriteString(", password=")
	builder.WriteString(m.Password)
	builder.WriteString(", comment=")
	builder.WriteString(m.Comment)
	builder.WriteString(", command=")
	builder.WriteString(m.Command)
	builder.WriteString(", creator=")
	builder.WriteString(m.Creator)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ctime=")
	builder.WriteString(m.Ctime.Format(time.ANSIC))
	builder.WriteString(", mtime=")
	builder.WriteString(m.Mtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Machines is a parsable slice of Machine.
type Machines []*Machine

func (m Machines) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
