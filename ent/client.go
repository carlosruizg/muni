// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/carlosruizg/muni/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/labellingtaskresponse"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Expert is the client for interacting with the Expert builders.
	Expert *ExpertClient
	// LabellingTask is the client for interacting with the LabellingTask builders.
	LabellingTask *LabellingTaskClient
	// LabellingTaskResponse is the client for interacting with the LabellingTaskResponse builders.
	LabellingTaskResponse *LabellingTaskResponseClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Expert = NewExpertClient(c.config)
	c.LabellingTask = NewLabellingTaskClient(c.config)
	c.LabellingTaskResponse = NewLabellingTaskResponseClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                   ctx,
		config:                cfg,
		Expert:                NewExpertClient(cfg),
		LabellingTask:         NewLabellingTaskClient(cfg),
		LabellingTaskResponse: NewLabellingTaskResponseClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                   ctx,
		config:                cfg,
		Expert:                NewExpertClient(cfg),
		LabellingTask:         NewLabellingTaskClient(cfg),
		LabellingTaskResponse: NewLabellingTaskResponseClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Expert.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Expert.Use(hooks...)
	c.LabellingTask.Use(hooks...)
	c.LabellingTaskResponse.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Expert.Intercept(interceptors...)
	c.LabellingTask.Intercept(interceptors...)
	c.LabellingTaskResponse.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ExpertMutation:
		return c.Expert.mutate(ctx, m)
	case *LabellingTaskMutation:
		return c.LabellingTask.mutate(ctx, m)
	case *LabellingTaskResponseMutation:
		return c.LabellingTaskResponse.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ExpertClient is a client for the Expert schema.
type ExpertClient struct {
	config
}

// NewExpertClient returns a client for the Expert from the given config.
func NewExpertClient(c config) *ExpertClient {
	return &ExpertClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `expert.Hooks(f(g(h())))`.
func (c *ExpertClient) Use(hooks ...Hook) {
	c.hooks.Expert = append(c.hooks.Expert, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `expert.Intercept(f(g(h())))`.
func (c *ExpertClient) Intercept(interceptors ...Interceptor) {
	c.inters.Expert = append(c.inters.Expert, interceptors...)
}

// Create returns a builder for creating a Expert entity.
func (c *ExpertClient) Create() *ExpertCreate {
	mutation := newExpertMutation(c.config, OpCreate)
	return &ExpertCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Expert entities.
func (c *ExpertClient) CreateBulk(builders ...*ExpertCreate) *ExpertCreateBulk {
	return &ExpertCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Expert.
func (c *ExpertClient) Update() *ExpertUpdate {
	mutation := newExpertMutation(c.config, OpUpdate)
	return &ExpertUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExpertClient) UpdateOne(e *Expert) *ExpertUpdateOne {
	mutation := newExpertMutation(c.config, OpUpdateOne, withExpert(e))
	return &ExpertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExpertClient) UpdateOneID(id int) *ExpertUpdateOne {
	mutation := newExpertMutation(c.config, OpUpdateOne, withExpertID(id))
	return &ExpertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Expert.
func (c *ExpertClient) Delete() *ExpertDelete {
	mutation := newExpertMutation(c.config, OpDelete)
	return &ExpertDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ExpertClient) DeleteOne(e *Expert) *ExpertDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ExpertClient) DeleteOneID(id int) *ExpertDeleteOne {
	builder := c.Delete().Where(expert.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExpertDeleteOne{builder}
}

// Query returns a query builder for Expert.
func (c *ExpertClient) Query() *ExpertQuery {
	return &ExpertQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeExpert},
		inters: c.Interceptors(),
	}
}

// Get returns a Expert entity by its id.
func (c *ExpertClient) Get(ctx context.Context, id int) (*Expert, error) {
	return c.Query().Where(expert.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExpertClient) GetX(ctx context.Context, id int) *Expert {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ExpertClient) Hooks() []Hook {
	return c.hooks.Expert
}

// Interceptors returns the client interceptors.
func (c *ExpertClient) Interceptors() []Interceptor {
	return c.inters.Expert
}

func (c *ExpertClient) mutate(ctx context.Context, m *ExpertMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ExpertCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ExpertUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ExpertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ExpertDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Expert mutation op: %q", m.Op())
	}
}

// LabellingTaskClient is a client for the LabellingTask schema.
type LabellingTaskClient struct {
	config
}

// NewLabellingTaskClient returns a client for the LabellingTask from the given config.
func NewLabellingTaskClient(c config) *LabellingTaskClient {
	return &LabellingTaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `labellingtask.Hooks(f(g(h())))`.
func (c *LabellingTaskClient) Use(hooks ...Hook) {
	c.hooks.LabellingTask = append(c.hooks.LabellingTask, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `labellingtask.Intercept(f(g(h())))`.
func (c *LabellingTaskClient) Intercept(interceptors ...Interceptor) {
	c.inters.LabellingTask = append(c.inters.LabellingTask, interceptors...)
}

// Create returns a builder for creating a LabellingTask entity.
func (c *LabellingTaskClient) Create() *LabellingTaskCreate {
	mutation := newLabellingTaskMutation(c.config, OpCreate)
	return &LabellingTaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LabellingTask entities.
func (c *LabellingTaskClient) CreateBulk(builders ...*LabellingTaskCreate) *LabellingTaskCreateBulk {
	return &LabellingTaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LabellingTask.
func (c *LabellingTaskClient) Update() *LabellingTaskUpdate {
	mutation := newLabellingTaskMutation(c.config, OpUpdate)
	return &LabellingTaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LabellingTaskClient) UpdateOne(lt *LabellingTask) *LabellingTaskUpdateOne {
	mutation := newLabellingTaskMutation(c.config, OpUpdateOne, withLabellingTask(lt))
	return &LabellingTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LabellingTaskClient) UpdateOneID(id int) *LabellingTaskUpdateOne {
	mutation := newLabellingTaskMutation(c.config, OpUpdateOne, withLabellingTaskID(id))
	return &LabellingTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LabellingTask.
func (c *LabellingTaskClient) Delete() *LabellingTaskDelete {
	mutation := newLabellingTaskMutation(c.config, OpDelete)
	return &LabellingTaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LabellingTaskClient) DeleteOne(lt *LabellingTask) *LabellingTaskDeleteOne {
	return c.DeleteOneID(lt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LabellingTaskClient) DeleteOneID(id int) *LabellingTaskDeleteOne {
	builder := c.Delete().Where(labellingtask.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LabellingTaskDeleteOne{builder}
}

// Query returns a query builder for LabellingTask.
func (c *LabellingTaskClient) Query() *LabellingTaskQuery {
	return &LabellingTaskQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLabellingTask},
		inters: c.Interceptors(),
	}
}

// Get returns a LabellingTask entity by its id.
func (c *LabellingTaskClient) Get(ctx context.Context, id int) (*LabellingTask, error) {
	return c.Query().Where(labellingtask.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LabellingTaskClient) GetX(ctx context.Context, id int) *LabellingTask {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResponses queries the responses edge of a LabellingTask.
func (c *LabellingTaskClient) QueryResponses(lt *LabellingTask) *LabellingTaskResponseQuery {
	query := (&LabellingTaskResponseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := lt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(labellingtask.Table, labellingtask.FieldID, id),
			sqlgraph.To(labellingtaskresponse.Table, labellingtaskresponse.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, labellingtask.ResponsesTable, labellingtask.ResponsesColumn),
		)
		fromV = sqlgraph.Neighbors(lt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LabellingTaskClient) Hooks() []Hook {
	return c.hooks.LabellingTask
}

// Interceptors returns the client interceptors.
func (c *LabellingTaskClient) Interceptors() []Interceptor {
	return c.inters.LabellingTask
}

func (c *LabellingTaskClient) mutate(ctx context.Context, m *LabellingTaskMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LabellingTaskCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LabellingTaskUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LabellingTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LabellingTaskDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LabellingTask mutation op: %q", m.Op())
	}
}

// LabellingTaskResponseClient is a client for the LabellingTaskResponse schema.
type LabellingTaskResponseClient struct {
	config
}

// NewLabellingTaskResponseClient returns a client for the LabellingTaskResponse from the given config.
func NewLabellingTaskResponseClient(c config) *LabellingTaskResponseClient {
	return &LabellingTaskResponseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `labellingtaskresponse.Hooks(f(g(h())))`.
func (c *LabellingTaskResponseClient) Use(hooks ...Hook) {
	c.hooks.LabellingTaskResponse = append(c.hooks.LabellingTaskResponse, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `labellingtaskresponse.Intercept(f(g(h())))`.
func (c *LabellingTaskResponseClient) Intercept(interceptors ...Interceptor) {
	c.inters.LabellingTaskResponse = append(c.inters.LabellingTaskResponse, interceptors...)
}

// Create returns a builder for creating a LabellingTaskResponse entity.
func (c *LabellingTaskResponseClient) Create() *LabellingTaskResponseCreate {
	mutation := newLabellingTaskResponseMutation(c.config, OpCreate)
	return &LabellingTaskResponseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LabellingTaskResponse entities.
func (c *LabellingTaskResponseClient) CreateBulk(builders ...*LabellingTaskResponseCreate) *LabellingTaskResponseCreateBulk {
	return &LabellingTaskResponseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LabellingTaskResponse.
func (c *LabellingTaskResponseClient) Update() *LabellingTaskResponseUpdate {
	mutation := newLabellingTaskResponseMutation(c.config, OpUpdate)
	return &LabellingTaskResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LabellingTaskResponseClient) UpdateOne(ltr *LabellingTaskResponse) *LabellingTaskResponseUpdateOne {
	mutation := newLabellingTaskResponseMutation(c.config, OpUpdateOne, withLabellingTaskResponse(ltr))
	return &LabellingTaskResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LabellingTaskResponseClient) UpdateOneID(id int) *LabellingTaskResponseUpdateOne {
	mutation := newLabellingTaskResponseMutation(c.config, OpUpdateOne, withLabellingTaskResponseID(id))
	return &LabellingTaskResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LabellingTaskResponse.
func (c *LabellingTaskResponseClient) Delete() *LabellingTaskResponseDelete {
	mutation := newLabellingTaskResponseMutation(c.config, OpDelete)
	return &LabellingTaskResponseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LabellingTaskResponseClient) DeleteOne(ltr *LabellingTaskResponse) *LabellingTaskResponseDeleteOne {
	return c.DeleteOneID(ltr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LabellingTaskResponseClient) DeleteOneID(id int) *LabellingTaskResponseDeleteOne {
	builder := c.Delete().Where(labellingtaskresponse.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LabellingTaskResponseDeleteOne{builder}
}

// Query returns a query builder for LabellingTaskResponse.
func (c *LabellingTaskResponseClient) Query() *LabellingTaskResponseQuery {
	return &LabellingTaskResponseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLabellingTaskResponse},
		inters: c.Interceptors(),
	}
}

// Get returns a LabellingTaskResponse entity by its id.
func (c *LabellingTaskResponseClient) Get(ctx context.Context, id int) (*LabellingTaskResponse, error) {
	return c.Query().Where(labellingtaskresponse.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LabellingTaskResponseClient) GetX(ctx context.Context, id int) *LabellingTaskResponse {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTask queries the task edge of a LabellingTaskResponse.
func (c *LabellingTaskResponseClient) QueryTask(ltr *LabellingTaskResponse) *LabellingTaskQuery {
	query := (&LabellingTaskClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ltr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(labellingtaskresponse.Table, labellingtaskresponse.FieldID, id),
			sqlgraph.To(labellingtask.Table, labellingtask.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, labellingtaskresponse.TaskTable, labellingtaskresponse.TaskColumn),
		)
		fromV = sqlgraph.Neighbors(ltr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LabellingTaskResponseClient) Hooks() []Hook {
	return c.hooks.LabellingTaskResponse
}

// Interceptors returns the client interceptors.
func (c *LabellingTaskResponseClient) Interceptors() []Interceptor {
	return c.inters.LabellingTaskResponse
}

func (c *LabellingTaskResponseClient) mutate(ctx context.Context, m *LabellingTaskResponseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LabellingTaskResponseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LabellingTaskResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LabellingTaskResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LabellingTaskResponseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LabellingTaskResponse mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Expert, LabellingTask, LabellingTaskResponse []ent.Hook
	}
	inters struct {
		Expert, LabellingTask, LabellingTaskResponse []ent.Interceptor
	}
)
