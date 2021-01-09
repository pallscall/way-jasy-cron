// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"graduate/cron/internal/model/ent/job"
	"graduate/cron/internal/model/ent/predicate"
	"sync"
	"time"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeJob = "Job"
)

// JobMutation represents an operation that mutate the Jobs
// nodes in the graph.
type JobMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	creator       *string
	url           *string
	spec          *string
	comment       *string
	updater       *string
	method        *string
	body          *string
	header        *string
	stoppable     *int
	addstoppable  *int
	status        *int
	addstatus     *int
	ctime         *time.Time
	mtime         *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Job, error)
	predicates    []predicate.Job
}

var _ ent.Mutation = (*JobMutation)(nil)

// jobOption allows to manage the mutation configuration using functional options.
type jobOption func(*JobMutation)

// newJobMutation creates new mutation for Job.
func newJobMutation(c config, op Op, opts ...jobOption) *JobMutation {
	m := &JobMutation{
		config:        c,
		op:            op,
		typ:           TypeJob,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withJobID sets the id field of the mutation.
func withJobID(id int) jobOption {
	return func(m *JobMutation) {
		var (
			err   error
			once  sync.Once
			value *Job
		)
		m.oldValue = func(ctx context.Context) (*Job, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Job.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withJob sets the old Job of the mutation.
func withJob(node *Job) jobOption {
	return func(m *JobMutation) {
		m.oldValue = func(context.Context) (*Job, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m JobMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m JobMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Job creation.
func (m *JobMutation) SetID(id int) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *JobMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *JobMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *JobMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *JobMutation) ResetName() {
	m.name = nil
}

// SetCreator sets the creator field.
func (m *JobMutation) SetCreator(s string) {
	m.creator = &s
}

// Creator returns the creator value in the mutation.
func (m *JobMutation) Creator() (r string, exists bool) {
	v := m.creator
	if v == nil {
		return
	}
	return *v, true
}

// OldCreator returns the old creator value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldCreator(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreator is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreator requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreator: %w", err)
	}
	return oldValue.Creator, nil
}

// ResetCreator reset all changes of the "creator" field.
func (m *JobMutation) ResetCreator() {
	m.creator = nil
}

// SetURL sets the url field.
func (m *JobMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the url value in the mutation.
func (m *JobMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old url value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL reset all changes of the "url" field.
func (m *JobMutation) ResetURL() {
	m.url = nil
}

// SetSpec sets the spec field.
func (m *JobMutation) SetSpec(s string) {
	m.spec = &s
}

// Spec returns the spec value in the mutation.
func (m *JobMutation) Spec() (r string, exists bool) {
	v := m.spec
	if v == nil {
		return
	}
	return *v, true
}

// OldSpec returns the old spec value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldSpec(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSpec is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSpec requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSpec: %w", err)
	}
	return oldValue.Spec, nil
}

// ResetSpec reset all changes of the "spec" field.
func (m *JobMutation) ResetSpec() {
	m.spec = nil
}

// SetComment sets the comment field.
func (m *JobMutation) SetComment(s string) {
	m.comment = &s
}

// Comment returns the comment value in the mutation.
func (m *JobMutation) Comment() (r string, exists bool) {
	v := m.comment
	if v == nil {
		return
	}
	return *v, true
}

// OldComment returns the old comment value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldComment(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldComment is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldComment requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldComment: %w", err)
	}
	return oldValue.Comment, nil
}

// ResetComment reset all changes of the "comment" field.
func (m *JobMutation) ResetComment() {
	m.comment = nil
}

// SetUpdater sets the updater field.
func (m *JobMutation) SetUpdater(s string) {
	m.updater = &s
}

// Updater returns the updater value in the mutation.
func (m *JobMutation) Updater() (r string, exists bool) {
	v := m.updater
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdater returns the old updater value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldUpdater(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdater is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdater requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdater: %w", err)
	}
	return oldValue.Updater, nil
}

// ResetUpdater reset all changes of the "updater" field.
func (m *JobMutation) ResetUpdater() {
	m.updater = nil
}

// SetMethod sets the method field.
func (m *JobMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the method value in the mutation.
func (m *JobMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old method value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMethod is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ResetMethod reset all changes of the "method" field.
func (m *JobMutation) ResetMethod() {
	m.method = nil
}

// SetBody sets the body field.
func (m *JobMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the body value in the mutation.
func (m *JobMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old body value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBody is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody reset all changes of the "body" field.
func (m *JobMutation) ResetBody() {
	m.body = nil
}

// SetHeader sets the header field.
func (m *JobMutation) SetHeader(s string) {
	m.header = &s
}

// Header returns the header value in the mutation.
func (m *JobMutation) Header() (r string, exists bool) {
	v := m.header
	if v == nil {
		return
	}
	return *v, true
}

// OldHeader returns the old header value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldHeader(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldHeader is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldHeader requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHeader: %w", err)
	}
	return oldValue.Header, nil
}

// ResetHeader reset all changes of the "header" field.
func (m *JobMutation) ResetHeader() {
	m.header = nil
}

// SetStoppable sets the stoppable field.
func (m *JobMutation) SetStoppable(i int) {
	m.stoppable = &i
	m.addstoppable = nil
}

// Stoppable returns the stoppable value in the mutation.
func (m *JobMutation) Stoppable() (r int, exists bool) {
	v := m.stoppable
	if v == nil {
		return
	}
	return *v, true
}

// OldStoppable returns the old stoppable value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldStoppable(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStoppable is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStoppable requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStoppable: %w", err)
	}
	return oldValue.Stoppable, nil
}

// AddStoppable adds i to stoppable.
func (m *JobMutation) AddStoppable(i int) {
	if m.addstoppable != nil {
		*m.addstoppable += i
	} else {
		m.addstoppable = &i
	}
}

// AddedStoppable returns the value that was added to the stoppable field in this mutation.
func (m *JobMutation) AddedStoppable() (r int, exists bool) {
	v := m.addstoppable
	if v == nil {
		return
	}
	return *v, true
}

// ResetStoppable reset all changes of the "stoppable" field.
func (m *JobMutation) ResetStoppable() {
	m.stoppable = nil
	m.addstoppable = nil
}

// SetStatus sets the status field.
func (m *JobMutation) SetStatus(i int) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the status value in the mutation.
func (m *JobMutation) Status() (r int, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old status value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldStatus(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds i to status.
func (m *JobMutation) AddStatus(i int) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the status field in this mutation.
func (m *JobMutation) AddedStatus() (r int, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus reset all changes of the "status" field.
func (m *JobMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// SetCtime sets the ctime field.
func (m *JobMutation) SetCtime(t time.Time) {
	m.ctime = &t
}

// Ctime returns the ctime value in the mutation.
func (m *JobMutation) Ctime() (r time.Time, exists bool) {
	v := m.ctime
	if v == nil {
		return
	}
	return *v, true
}

// OldCtime returns the old ctime value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldCtime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCtime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCtime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCtime: %w", err)
	}
	return oldValue.Ctime, nil
}

// ClearCtime clears the value of ctime.
func (m *JobMutation) ClearCtime() {
	m.ctime = nil
	m.clearedFields[job.FieldCtime] = struct{}{}
}

// CtimeCleared returns if the field ctime was cleared in this mutation.
func (m *JobMutation) CtimeCleared() bool {
	_, ok := m.clearedFields[job.FieldCtime]
	return ok
}

// ResetCtime reset all changes of the "ctime" field.
func (m *JobMutation) ResetCtime() {
	m.ctime = nil
	delete(m.clearedFields, job.FieldCtime)
}

// SetMtime sets the mtime field.
func (m *JobMutation) SetMtime(t time.Time) {
	m.mtime = &t
}

// Mtime returns the mtime value in the mutation.
func (m *JobMutation) Mtime() (r time.Time, exists bool) {
	v := m.mtime
	if v == nil {
		return
	}
	return *v, true
}

// OldMtime returns the old mtime value of the Job.
// If the Job object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *JobMutation) OldMtime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMtime is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMtime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMtime: %w", err)
	}
	return oldValue.Mtime, nil
}

// ClearMtime clears the value of mtime.
func (m *JobMutation) ClearMtime() {
	m.mtime = nil
	m.clearedFields[job.FieldMtime] = struct{}{}
}

// MtimeCleared returns if the field mtime was cleared in this mutation.
func (m *JobMutation) MtimeCleared() bool {
	_, ok := m.clearedFields[job.FieldMtime]
	return ok
}

// ResetMtime reset all changes of the "mtime" field.
func (m *JobMutation) ResetMtime() {
	m.mtime = nil
	delete(m.clearedFields, job.FieldMtime)
}

// Op returns the operation name.
func (m *JobMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Job).
func (m *JobMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *JobMutation) Fields() []string {
	fields := make([]string, 0, 13)
	if m.name != nil {
		fields = append(fields, job.FieldName)
	}
	if m.creator != nil {
		fields = append(fields, job.FieldCreator)
	}
	if m.url != nil {
		fields = append(fields, job.FieldURL)
	}
	if m.spec != nil {
		fields = append(fields, job.FieldSpec)
	}
	if m.comment != nil {
		fields = append(fields, job.FieldComment)
	}
	if m.updater != nil {
		fields = append(fields, job.FieldUpdater)
	}
	if m.method != nil {
		fields = append(fields, job.FieldMethod)
	}
	if m.body != nil {
		fields = append(fields, job.FieldBody)
	}
	if m.header != nil {
		fields = append(fields, job.FieldHeader)
	}
	if m.stoppable != nil {
		fields = append(fields, job.FieldStoppable)
	}
	if m.status != nil {
		fields = append(fields, job.FieldStatus)
	}
	if m.ctime != nil {
		fields = append(fields, job.FieldCtime)
	}
	if m.mtime != nil {
		fields = append(fields, job.FieldMtime)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *JobMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case job.FieldName:
		return m.Name()
	case job.FieldCreator:
		return m.Creator()
	case job.FieldURL:
		return m.URL()
	case job.FieldSpec:
		return m.Spec()
	case job.FieldComment:
		return m.Comment()
	case job.FieldUpdater:
		return m.Updater()
	case job.FieldMethod:
		return m.Method()
	case job.FieldBody:
		return m.Body()
	case job.FieldHeader:
		return m.Header()
	case job.FieldStoppable:
		return m.Stoppable()
	case job.FieldStatus:
		return m.Status()
	case job.FieldCtime:
		return m.Ctime()
	case job.FieldMtime:
		return m.Mtime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *JobMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case job.FieldName:
		return m.OldName(ctx)
	case job.FieldCreator:
		return m.OldCreator(ctx)
	case job.FieldURL:
		return m.OldURL(ctx)
	case job.FieldSpec:
		return m.OldSpec(ctx)
	case job.FieldComment:
		return m.OldComment(ctx)
	case job.FieldUpdater:
		return m.OldUpdater(ctx)
	case job.FieldMethod:
		return m.OldMethod(ctx)
	case job.FieldBody:
		return m.OldBody(ctx)
	case job.FieldHeader:
		return m.OldHeader(ctx)
	case job.FieldStoppable:
		return m.OldStoppable(ctx)
	case job.FieldStatus:
		return m.OldStatus(ctx)
	case job.FieldCtime:
		return m.OldCtime(ctx)
	case job.FieldMtime:
		return m.OldMtime(ctx)
	}
	return nil, fmt.Errorf("unknown Job field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *JobMutation) SetField(name string, value ent.Value) error {
	switch name {
	case job.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case job.FieldCreator:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreator(v)
		return nil
	case job.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case job.FieldSpec:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSpec(v)
		return nil
	case job.FieldComment:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetComment(v)
		return nil
	case job.FieldUpdater:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdater(v)
		return nil
	case job.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	case job.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case job.FieldHeader:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHeader(v)
		return nil
	case job.FieldStoppable:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStoppable(v)
		return nil
	case job.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case job.FieldCtime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCtime(v)
		return nil
	case job.FieldMtime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMtime(v)
		return nil
	}
	return fmt.Errorf("unknown Job field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *JobMutation) AddedFields() []string {
	var fields []string
	if m.addstoppable != nil {
		fields = append(fields, job.FieldStoppable)
	}
	if m.addstatus != nil {
		fields = append(fields, job.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *JobMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case job.FieldStoppable:
		return m.AddedStoppable()
	case job.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *JobMutation) AddField(name string, value ent.Value) error {
	switch name {
	case job.FieldStoppable:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStoppable(v)
		return nil
	case job.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Job numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *JobMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(job.FieldCtime) {
		fields = append(fields, job.FieldCtime)
	}
	if m.FieldCleared(job.FieldMtime) {
		fields = append(fields, job.FieldMtime)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *JobMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *JobMutation) ClearField(name string) error {
	switch name {
	case job.FieldCtime:
		m.ClearCtime()
		return nil
	case job.FieldMtime:
		m.ClearMtime()
		return nil
	}
	return fmt.Errorf("unknown Job nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *JobMutation) ResetField(name string) error {
	switch name {
	case job.FieldName:
		m.ResetName()
		return nil
	case job.FieldCreator:
		m.ResetCreator()
		return nil
	case job.FieldURL:
		m.ResetURL()
		return nil
	case job.FieldSpec:
		m.ResetSpec()
		return nil
	case job.FieldComment:
		m.ResetComment()
		return nil
	case job.FieldUpdater:
		m.ResetUpdater()
		return nil
	case job.FieldMethod:
		m.ResetMethod()
		return nil
	case job.FieldBody:
		m.ResetBody()
		return nil
	case job.FieldHeader:
		m.ResetHeader()
		return nil
	case job.FieldStoppable:
		m.ResetStoppable()
		return nil
	case job.FieldStatus:
		m.ResetStatus()
		return nil
	case job.FieldCtime:
		m.ResetCtime()
		return nil
	case job.FieldMtime:
		m.ResetMtime()
		return nil
	}
	return fmt.Errorf("unknown Job field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *JobMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *JobMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *JobMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *JobMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *JobMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *JobMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *JobMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Job unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *JobMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Job edge %s", name)
}