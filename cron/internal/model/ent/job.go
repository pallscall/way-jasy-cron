// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"way-jasy-cron/cron/internal/model/ent/job"

	"github.com/facebook/ent/dialect/sql"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Spec holds the value of the "spec" field.
	Spec string `json:"spec,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Updater holds the value of the "updater" field.
	Updater string `json:"updater,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Header holds the value of the "header" field.
	Header string `json:"header,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// Retry holds the value of the "retry" field.
	Retry int `json:"retry,omitempty"`
	// RetryTemp holds the value of the "retry_temp" field.
	RetryTemp int `json:"retry_temp,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Ctime holds the value of the "ctime" field.
	Ctime time.Time `json:"ctime,omitempty"`
	// Mtime holds the value of the "mtime" field.
	Mtime time.Time `json:"mtime,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldID, job.FieldCount, job.FieldRetry, job.FieldRetryTemp, job.FieldStatus:
			values[i] = &sql.NullInt64{}
		case job.FieldName, job.FieldCreator, job.FieldURL, job.FieldSpec, job.FieldComment, job.FieldUpdater, job.FieldMethod, job.FieldBody, job.FieldHeader:
			values[i] = &sql.NullString{}
		case job.FieldCtime, job.FieldMtime:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Job", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = int(value.Int64)
		case job.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				j.Name = value.String
			}
		case job.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				j.Creator = value.String
			}
		case job.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				j.URL = value.String
			}
		case job.FieldSpec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field spec", values[i])
			} else if value.Valid {
				j.Spec = value.String
			}
		case job.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				j.Comment = value.String
			}
		case job.FieldUpdater:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updater", values[i])
			} else if value.Valid {
				j.Updater = value.String
			}
		case job.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				j.Method = value.String
			}
		case job.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				j.Body = value.String
			}
		case job.FieldHeader:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field header", values[i])
			} else if value.Valid {
				j.Header = value.String
			}
		case job.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				j.Count = int(value.Int64)
			}
		case job.FieldRetry:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field retry", values[i])
			} else if value.Valid {
				j.Retry = int(value.Int64)
			}
		case job.FieldRetryTemp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field retry_temp", values[i])
			} else if value.Valid {
				j.RetryTemp = int(value.Int64)
			}
		case job.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				j.Status = int(value.Int64)
			}
		case job.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				j.Ctime = value.Time
			}
		case job.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				j.Mtime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return (&JobClient{config: j.config}).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v", j.ID))
	builder.WriteString(", name=")
	builder.WriteString(j.Name)
	builder.WriteString(", creator=")
	builder.WriteString(j.Creator)
	builder.WriteString(", url=")
	builder.WriteString(j.URL)
	builder.WriteString(", spec=")
	builder.WriteString(j.Spec)
	builder.WriteString(", comment=")
	builder.WriteString(j.Comment)
	builder.WriteString(", updater=")
	builder.WriteString(j.Updater)
	builder.WriteString(", method=")
	builder.WriteString(j.Method)
	builder.WriteString(", body=")
	builder.WriteString(j.Body)
	builder.WriteString(", header=")
	builder.WriteString(j.Header)
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", j.Count))
	builder.WriteString(", retry=")
	builder.WriteString(fmt.Sprintf("%v", j.Retry))
	builder.WriteString(", retry_temp=")
	builder.WriteString(fmt.Sprintf("%v", j.RetryTemp))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", j.Status))
	builder.WriteString(", ctime=")
	builder.WriteString(j.Ctime.Format(time.ANSIC))
	builder.WriteString(", mtime=")
	builder.WriteString(j.Mtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job

func (j Jobs) config(cfg config) {
	for _i := range j {
		j[_i].config = cfg
	}
}
