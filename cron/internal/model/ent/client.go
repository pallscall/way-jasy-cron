// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"way-jasy-cron/cron/internal/model/ent/migrate"

	"way-jasy-cron/cron/internal/model/ent/job"
	"way-jasy-cron/cron/internal/model/ent/machine"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// Machine is the client for interacting with the Machine builders.
	Machine *MachineClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Job = NewJobClient(c.config)
	c.Machine = NewMachineClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Job:     NewJobClient(cfg),
		Machine: NewMachineClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:  cfg,
		Job:     NewJobClient(cfg),
		Machine: NewMachineClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Job.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Job.Use(hooks...)
	c.Machine.Use(hooks...)
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `job.Hooks(f(g(h())))`.
func (c *JobClient) Use(hooks ...Hook) {
	c.hooks.Job = append(c.hooks.Job, hooks...)
}

// Create returns a create builder for Job.
func (c *JobClient) Create() *JobCreate {
	mutation := newJobMutation(c.config, OpCreate)
	return &JobCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Job entities.
func (c *JobClient) CreateBulk(builders ...*JobCreate) *JobCreateBulk {
	return &JobCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	mutation := newJobMutation(c.config, OpUpdate)
	return &JobUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJob(j))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id int) *JobUpdateOne {
	mutation := newJobMutation(c.config, OpUpdateOne, withJobID(id))
	return &JobUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	mutation := newJobMutation(c.config, OpDelete)
	return &JobDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *JobClient) DeleteOneID(id int) *JobDeleteOne {
	builder := c.Delete().Where(job.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &JobDeleteOne{builder}
}

// Query returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{config: c.config}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id int) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id int) *Job {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *JobClient) Hooks() []Hook {
	return c.hooks.Job
}

// MachineClient is a client for the Machine schema.
type MachineClient struct {
	config
}

// NewMachineClient returns a client for the Machine from the given config.
func NewMachineClient(c config) *MachineClient {
	return &MachineClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `machine.Hooks(f(g(h())))`.
func (c *MachineClient) Use(hooks ...Hook) {
	c.hooks.Machine = append(c.hooks.Machine, hooks...)
}

// Create returns a create builder for Machine.
func (c *MachineClient) Create() *MachineCreate {
	mutation := newMachineMutation(c.config, OpCreate)
	return &MachineCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Machine entities.
func (c *MachineClient) CreateBulk(builders ...*MachineCreate) *MachineCreateBulk {
	return &MachineCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Machine.
func (c *MachineClient) Update() *MachineUpdate {
	mutation := newMachineMutation(c.config, OpUpdate)
	return &MachineUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MachineClient) UpdateOne(m *Machine) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachine(m))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MachineClient) UpdateOneID(id int) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachineID(id))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Machine.
func (c *MachineClient) Delete() *MachineDelete {
	mutation := newMachineMutation(c.config, OpDelete)
	return &MachineDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MachineClient) DeleteOne(m *Machine) *MachineDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MachineClient) DeleteOneID(id int) *MachineDeleteOne {
	builder := c.Delete().Where(machine.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MachineDeleteOne{builder}
}

// Query returns a query builder for Machine.
func (c *MachineClient) Query() *MachineQuery {
	return &MachineQuery{config: c.config}
}

// Get returns a Machine entity by its id.
func (c *MachineClient) Get(ctx context.Context, id int) (*Machine, error) {
	return c.Query().Where(machine.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MachineClient) GetX(ctx context.Context, id int) *Machine {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MachineClient) Hooks() []Hook {
	return c.hooks.Machine
}
