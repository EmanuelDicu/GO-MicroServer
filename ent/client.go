// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"tema-sprc-go/ent/migrate"

	"tema-sprc-go/ent/city"
	"tema-sprc-go/ent/country"
	"tema-sprc-go/ent/temperature"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// City is the client for interacting with the City builders.
	City *CityClient
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
	// Temperature is the client for interacting with the Temperature builders.
	Temperature *TemperatureClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.City = NewCityClient(c.config)
	c.Country = NewCountryClient(c.config)
	c.Temperature = NewTemperatureClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		City:        NewCityClient(cfg),
		Country:     NewCountryClient(cfg),
		Temperature: NewTemperatureClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		City:        NewCityClient(cfg),
		Country:     NewCountryClient(cfg),
		Temperature: NewTemperatureClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		City.
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
	c.City.Use(hooks...)
	c.Country.Use(hooks...)
	c.Temperature.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.City.Intercept(interceptors...)
	c.Country.Intercept(interceptors...)
	c.Temperature.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CityMutation:
		return c.City.mutate(ctx, m)
	case *CountryMutation:
		return c.Country.mutate(ctx, m)
	case *TemperatureMutation:
		return c.Temperature.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CityClient is a client for the City schema.
type CityClient struct {
	config
}

// NewCityClient returns a client for the City from the given config.
func NewCityClient(c config) *CityClient {
	return &CityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `city.Hooks(f(g(h())))`.
func (c *CityClient) Use(hooks ...Hook) {
	c.hooks.City = append(c.hooks.City, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `city.Intercept(f(g(h())))`.
func (c *CityClient) Intercept(interceptors ...Interceptor) {
	c.inters.City = append(c.inters.City, interceptors...)
}

// Create returns a builder for creating a City entity.
func (c *CityClient) Create() *CityCreate {
	mutation := newCityMutation(c.config, OpCreate)
	return &CityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of City entities.
func (c *CityClient) CreateBulk(builders ...*CityCreate) *CityCreateBulk {
	return &CityCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CityClient) MapCreateBulk(slice any, setFunc func(*CityCreate, int)) *CityCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CityCreateBulk{err: fmt.Errorf("calling to CityClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CityCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for City.
func (c *CityClient) Update() *CityUpdate {
	mutation := newCityMutation(c.config, OpUpdate)
	return &CityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CityClient) UpdateOne(ci *City) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCity(ci))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CityClient) UpdateOneID(id int) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCityID(id))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for City.
func (c *CityClient) Delete() *CityDelete {
	mutation := newCityMutation(c.config, OpDelete)
	return &CityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CityClient) DeleteOne(ci *City) *CityDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CityClient) DeleteOneID(id int) *CityDeleteOne {
	builder := c.Delete().Where(city.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CityDeleteOne{builder}
}

// Query returns a query builder for City.
func (c *CityClient) Query() *CityQuery {
	return &CityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCity},
		inters: c.Interceptors(),
	}
}

// Get returns a City entity by its id.
func (c *CityClient) Get(ctx context.Context, id int) (*City, error) {
	return c.Query().Where(city.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CityClient) GetX(ctx context.Context, id int) *City {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTemperatures queries the Temperatures edge of a City.
func (c *CityClient) QueryTemperatures(ci *City) *TemperatureQuery {
	query := (&TemperatureClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(city.Table, city.FieldID, id),
			sqlgraph.To(temperature.Table, temperature.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, city.TemperaturesTable, city.TemperaturesColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CityClient) Hooks() []Hook {
	return c.hooks.City
}

// Interceptors returns the client interceptors.
func (c *CityClient) Interceptors() []Interceptor {
	return c.inters.City
}

func (c *CityClient) mutate(ctx context.Context, m *CityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown City mutation op: %q", m.Op())
	}
}

// CountryClient is a client for the Country schema.
type CountryClient struct {
	config
}

// NewCountryClient returns a client for the Country from the given config.
func NewCountryClient(c config) *CountryClient {
	return &CountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country.Hooks(f(g(h())))`.
func (c *CountryClient) Use(hooks ...Hook) {
	c.hooks.Country = append(c.hooks.Country, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `country.Intercept(f(g(h())))`.
func (c *CountryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Country = append(c.inters.Country, interceptors...)
}

// Create returns a builder for creating a Country entity.
func (c *CountryClient) Create() *CountryCreate {
	mutation := newCountryMutation(c.config, OpCreate)
	return &CountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country entities.
func (c *CountryClient) CreateBulk(builders ...*CountryCreate) *CountryCreateBulk {
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CountryClient) MapCreateBulk(slice any, setFunc func(*CountryCreate, int)) *CountryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CountryCreateBulk{err: fmt.Errorf("calling to CountryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CountryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country.
func (c *CountryClient) Update() *CountryUpdate {
	mutation := newCountryMutation(c.config, OpUpdate)
	return &CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountryClient) UpdateOne(co *Country) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountry(co))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountryClient) UpdateOneID(id int) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountryID(id))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country.
func (c *CountryClient) Delete() *CountryDelete {
	mutation := newCountryMutation(c.config, OpDelete)
	return &CountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CountryClient) DeleteOne(co *Country) *CountryDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CountryClient) DeleteOneID(id int) *CountryDeleteOne {
	builder := c.Delete().Where(country.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryDeleteOne{builder}
}

// Query returns a query builder for Country.
func (c *CountryClient) Query() *CountryQuery {
	return &CountryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCountry},
		inters: c.Interceptors(),
	}
}

// Get returns a Country entity by its id.
func (c *CountryClient) Get(ctx context.Context, id int) (*Country, error) {
	return c.Query().Where(country.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountryClient) GetX(ctx context.Context, id int) *Country {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCities queries the Cities edge of a Country.
func (c *CountryClient) QueryCities(co *Country) *CityQuery {
	query := (&CityClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, id),
			sqlgraph.To(city.Table, city.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, country.CitiesTable, country.CitiesColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CountryClient) Hooks() []Hook {
	return c.hooks.Country
}

// Interceptors returns the client interceptors.
func (c *CountryClient) Interceptors() []Interceptor {
	return c.inters.Country
}

func (c *CountryClient) mutate(ctx context.Context, m *CountryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CountryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CountryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Country mutation op: %q", m.Op())
	}
}

// TemperatureClient is a client for the Temperature schema.
type TemperatureClient struct {
	config
}

// NewTemperatureClient returns a client for the Temperature from the given config.
func NewTemperatureClient(c config) *TemperatureClient {
	return &TemperatureClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `temperature.Hooks(f(g(h())))`.
func (c *TemperatureClient) Use(hooks ...Hook) {
	c.hooks.Temperature = append(c.hooks.Temperature, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `temperature.Intercept(f(g(h())))`.
func (c *TemperatureClient) Intercept(interceptors ...Interceptor) {
	c.inters.Temperature = append(c.inters.Temperature, interceptors...)
}

// Create returns a builder for creating a Temperature entity.
func (c *TemperatureClient) Create() *TemperatureCreate {
	mutation := newTemperatureMutation(c.config, OpCreate)
	return &TemperatureCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Temperature entities.
func (c *TemperatureClient) CreateBulk(builders ...*TemperatureCreate) *TemperatureCreateBulk {
	return &TemperatureCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TemperatureClient) MapCreateBulk(slice any, setFunc func(*TemperatureCreate, int)) *TemperatureCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TemperatureCreateBulk{err: fmt.Errorf("calling to TemperatureClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TemperatureCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TemperatureCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Temperature.
func (c *TemperatureClient) Update() *TemperatureUpdate {
	mutation := newTemperatureMutation(c.config, OpUpdate)
	return &TemperatureUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TemperatureClient) UpdateOne(t *Temperature) *TemperatureUpdateOne {
	mutation := newTemperatureMutation(c.config, OpUpdateOne, withTemperature(t))
	return &TemperatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TemperatureClient) UpdateOneID(id int) *TemperatureUpdateOne {
	mutation := newTemperatureMutation(c.config, OpUpdateOne, withTemperatureID(id))
	return &TemperatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Temperature.
func (c *TemperatureClient) Delete() *TemperatureDelete {
	mutation := newTemperatureMutation(c.config, OpDelete)
	return &TemperatureDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TemperatureClient) DeleteOne(t *Temperature) *TemperatureDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TemperatureClient) DeleteOneID(id int) *TemperatureDeleteOne {
	builder := c.Delete().Where(temperature.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TemperatureDeleteOne{builder}
}

// Query returns a query builder for Temperature.
func (c *TemperatureClient) Query() *TemperatureQuery {
	return &TemperatureQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTemperature},
		inters: c.Interceptors(),
	}
}

// Get returns a Temperature entity by its id.
func (c *TemperatureClient) Get(ctx context.Context, id int) (*Temperature, error) {
	return c.Query().Where(temperature.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TemperatureClient) GetX(ctx context.Context, id int) *Temperature {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TemperatureClient) Hooks() []Hook {
	return c.hooks.Temperature
}

// Interceptors returns the client interceptors.
func (c *TemperatureClient) Interceptors() []Interceptor {
	return c.inters.Temperature
}

func (c *TemperatureClient) mutate(ctx context.Context, m *TemperatureMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TemperatureCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TemperatureUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TemperatureUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TemperatureDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Temperature mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		City, Country, Temperature []ent.Hook
	}
	inters struct {
		City, Country, Temperature []ent.Interceptor
	}
)
