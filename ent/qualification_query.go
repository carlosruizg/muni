// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/carlosruizg/muni/ent/expert"
	"github.com/carlosruizg/muni/ent/labellingtask"
	"github.com/carlosruizg/muni/ent/predicate"
	"github.com/carlosruizg/muni/ent/qualification"
)

// QualificationQuery is the builder for querying Qualification entities.
type QualificationQuery struct {
	config
	ctx              *QueryContext
	order            []qualification.OrderOption
	inters           []Interceptor
	predicates       []predicate.Qualification
	withTasks        *LabellingTaskQuery
	withExperts      *ExpertQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Qualification) error
	withNamedTasks   map[string]*LabellingTaskQuery
	withNamedExperts map[string]*ExpertQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QualificationQuery builder.
func (qq *QualificationQuery) Where(ps ...predicate.Qualification) *QualificationQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit the number of records to be returned by this query.
func (qq *QualificationQuery) Limit(limit int) *QualificationQuery {
	qq.ctx.Limit = &limit
	return qq
}

// Offset to start from.
func (qq *QualificationQuery) Offset(offset int) *QualificationQuery {
	qq.ctx.Offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QualificationQuery) Unique(unique bool) *QualificationQuery {
	qq.ctx.Unique = &unique
	return qq
}

// Order specifies how the records should be ordered.
func (qq *QualificationQuery) Order(o ...qualification.OrderOption) *QualificationQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryTasks chains the current query on the "tasks" edge.
func (qq *QualificationQuery) QueryTasks() *LabellingTaskQuery {
	query := (&LabellingTaskClient{config: qq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(qualification.Table, qualification.FieldID, selector),
			sqlgraph.To(labellingtask.Table, labellingtask.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, qualification.TasksTable, qualification.TasksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExperts chains the current query on the "experts" edge.
func (qq *QualificationQuery) QueryExperts() *ExpertQuery {
	query := (&ExpertClient{config: qq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(qualification.Table, qualification.FieldID, selector),
			sqlgraph.To(expert.Table, expert.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, qualification.ExpertsTable, qualification.ExpertsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Qualification entity from the query.
// Returns a *NotFoundError when no Qualification was found.
func (qq *QualificationQuery) First(ctx context.Context) (*Qualification, error) {
	nodes, err := qq.Limit(1).All(setContextOp(ctx, qq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{qualification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QualificationQuery) FirstX(ctx context.Context) *Qualification {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Qualification ID from the query.
// Returns a *NotFoundError when no Qualification ID was found.
func (qq *QualificationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(setContextOp(ctx, qq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{qualification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QualificationQuery) FirstIDX(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Qualification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Qualification entity is found.
// Returns a *NotFoundError when no Qualification entities are found.
func (qq *QualificationQuery) Only(ctx context.Context) (*Qualification, error) {
	nodes, err := qq.Limit(2).All(setContextOp(ctx, qq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{qualification.Label}
	default:
		return nil, &NotSingularError{qualification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QualificationQuery) OnlyX(ctx context.Context) *Qualification {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Qualification ID in the query.
// Returns a *NotSingularError when more than one Qualification ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QualificationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(setContextOp(ctx, qq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{qualification.Label}
	default:
		err = &NotSingularError{qualification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QualificationQuery) OnlyIDX(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Qualifications.
func (qq *QualificationQuery) All(ctx context.Context) ([]*Qualification, error) {
	ctx = setContextOp(ctx, qq.ctx, "All")
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Qualification, *QualificationQuery]()
	return withInterceptors[[]*Qualification](ctx, qq, qr, qq.inters)
}

// AllX is like All, but panics if an error occurs.
func (qq *QualificationQuery) AllX(ctx context.Context) []*Qualification {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Qualification IDs.
func (qq *QualificationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if qq.ctx.Unique == nil && qq.path != nil {
		qq.Unique(true)
	}
	ctx = setContextOp(ctx, qq.ctx, "IDs")
	if err = qq.Select(qualification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QualificationQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QualificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, qq.ctx, "Count")
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, qq, querierCount[*QualificationQuery](), qq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QualificationQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QualificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, qq.ctx, "Exist")
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QualificationQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QualificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QualificationQuery) Clone() *QualificationQuery {
	if qq == nil {
		return nil
	}
	return &QualificationQuery{
		config:      qq.config,
		ctx:         qq.ctx.Clone(),
		order:       append([]qualification.OrderOption{}, qq.order...),
		inters:      append([]Interceptor{}, qq.inters...),
		predicates:  append([]predicate.Qualification{}, qq.predicates...),
		withTasks:   qq.withTasks.Clone(),
		withExperts: qq.withExperts.Clone(),
		// clone intermediate query.
		sql:  qq.sql.Clone(),
		path: qq.path,
	}
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QualificationQuery) WithTasks(opts ...func(*LabellingTaskQuery)) *QualificationQuery {
	query := (&LabellingTaskClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	qq.withTasks = query
	return qq
}

// WithExperts tells the query-builder to eager-load the nodes that are connected to
// the "experts" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QualificationQuery) WithExperts(opts ...func(*ExpertQuery)) *QualificationQuery {
	query := (&ExpertClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	qq.withExperts = query
	return qq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value schema.QualificationValue `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Qualification.Query().
//		GroupBy(qualification.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QualificationQuery) GroupBy(field string, fields ...string) *QualificationGroupBy {
	qq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &QualificationGroupBy{build: qq}
	grbuild.flds = &qq.ctx.Fields
	grbuild.label = qualification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value schema.QualificationValue `json:"value,omitempty"`
//	}
//
//	client.Qualification.Query().
//		Select(qualification.FieldValue).
//		Scan(ctx, &v)
func (qq *QualificationQuery) Select(fields ...string) *QualificationSelect {
	qq.ctx.Fields = append(qq.ctx.Fields, fields...)
	sbuild := &QualificationSelect{QualificationQuery: qq}
	sbuild.label = qualification.Label
	sbuild.flds, sbuild.scan = &qq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a QualificationSelect configured with the given aggregations.
func (qq *QualificationQuery) Aggregate(fns ...AggregateFunc) *QualificationSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QualificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range qq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, qq); err != nil {
				return err
			}
		}
	}
	for _, f := range qq.ctx.Fields {
		if !qualification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QualificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Qualification, error) {
	var (
		nodes       = []*Qualification{}
		_spec       = qq.querySpec()
		loadedTypes = [2]bool{
			qq.withTasks != nil,
			qq.withExperts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Qualification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Qualification{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(qq.modifiers) > 0 {
		_spec.Modifiers = qq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qq.withTasks; query != nil {
		if err := qq.loadTasks(ctx, query, nodes,
			func(n *Qualification) { n.Edges.Tasks = []*LabellingTask{} },
			func(n *Qualification, e *LabellingTask) { n.Edges.Tasks = append(n.Edges.Tasks, e) }); err != nil {
			return nil, err
		}
	}
	if query := qq.withExperts; query != nil {
		if err := qq.loadExperts(ctx, query, nodes,
			func(n *Qualification) { n.Edges.Experts = []*Expert{} },
			func(n *Qualification, e *Expert) { n.Edges.Experts = append(n.Edges.Experts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range qq.withNamedTasks {
		if err := qq.loadTasks(ctx, query, nodes,
			func(n *Qualification) { n.appendNamedTasks(name) },
			func(n *Qualification, e *LabellingTask) { n.appendNamedTasks(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range qq.withNamedExperts {
		if err := qq.loadExperts(ctx, query, nodes,
			func(n *Qualification) { n.appendNamedExperts(name) },
			func(n *Qualification, e *Expert) { n.appendNamedExperts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range qq.loadTotal {
		if err := qq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qq *QualificationQuery) loadTasks(ctx context.Context, query *LabellingTaskQuery, nodes []*Qualification, init func(*Qualification), assign func(*Qualification, *LabellingTask)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Qualification)
	nids := make(map[int]map[*Qualification]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(qualification.TasksTable)
		s.Join(joinT).On(s.C(labellingtask.FieldID), joinT.C(qualification.TasksPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(qualification.TasksPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(qualification.TasksPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Qualification]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*LabellingTask](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tasks" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (qq *QualificationQuery) loadExperts(ctx context.Context, query *ExpertQuery, nodes []*Qualification, init func(*Qualification), assign func(*Qualification, *Expert)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Qualification)
	nids := make(map[int]map[*Qualification]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(qualification.ExpertsTable)
		s.Join(joinT).On(s.C(expert.FieldID), joinT.C(qualification.ExpertsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(qualification.ExpertsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(qualification.ExpertsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Qualification]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Expert](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "experts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (qq *QualificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	if len(qq.modifiers) > 0 {
		_spec.Modifiers = qq.modifiers
	}
	_spec.Node.Columns = qq.ctx.Fields
	if len(qq.ctx.Fields) > 0 {
		_spec.Unique = qq.ctx.Unique != nil && *qq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QualificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(qualification.Table, qualification.Columns, sqlgraph.NewFieldSpec(qualification.FieldID, field.TypeInt))
	_spec.From = qq.sql
	if unique := qq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if qq.path != nil {
		_spec.Unique = true
	}
	if fields := qq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, qualification.FieldID)
		for i := range fields {
			if fields[i] != qualification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QualificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(qualification.Table)
	columns := qq.ctx.Fields
	if len(columns) == 0 {
		columns = qualification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.ctx.Unique != nil && *qq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTasks tells the query-builder to eager-load the nodes that are connected to the "tasks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (qq *QualificationQuery) WithNamedTasks(name string, opts ...func(*LabellingTaskQuery)) *QualificationQuery {
	query := (&LabellingTaskClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if qq.withNamedTasks == nil {
		qq.withNamedTasks = make(map[string]*LabellingTaskQuery)
	}
	qq.withNamedTasks[name] = query
	return qq
}

// WithNamedExperts tells the query-builder to eager-load the nodes that are connected to the "experts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (qq *QualificationQuery) WithNamedExperts(name string, opts ...func(*ExpertQuery)) *QualificationQuery {
	query := (&ExpertClient{config: qq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if qq.withNamedExperts == nil {
		qq.withNamedExperts = make(map[string]*ExpertQuery)
	}
	qq.withNamedExperts[name] = query
	return qq
}

// QualificationGroupBy is the group-by builder for Qualification entities.
type QualificationGroupBy struct {
	selector
	build *QualificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QualificationGroupBy) Aggregate(fns ...AggregateFunc) *QualificationGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the selector query and scans the result into the given value.
func (qgb *QualificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qgb.build.ctx, "GroupBy")
	if err := qgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QualificationQuery, *QualificationGroupBy](ctx, qgb.build, qgb, qgb.build.inters, v)
}

func (qgb *QualificationGroupBy) sqlScan(ctx context.Context, root *QualificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*qgb.flds)+len(qgb.fns))
		for _, f := range *qgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*qgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// QualificationSelect is the builder for selecting fields of Qualification entities.
type QualificationSelect struct {
	*QualificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QualificationSelect) Aggregate(fns ...AggregateFunc) *QualificationSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QualificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, qs.ctx, "Select")
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*QualificationQuery, *QualificationSelect](ctx, qs.QualificationQuery, qs, qs.inters, v)
}

func (qs *QualificationSelect) sqlScan(ctx context.Context, root *QualificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}