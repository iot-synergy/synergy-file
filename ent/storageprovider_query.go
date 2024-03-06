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
	"github.com/iot-synergy/synergy-file/ent/cloudfile"
	"github.com/iot-synergy/synergy-file/ent/predicate"
	"github.com/iot-synergy/synergy-file/ent/storageprovider"
)

// StorageProviderQuery is the builder for querying StorageProvider entities.
type StorageProviderQuery struct {
	config
	ctx            *QueryContext
	order          []storageprovider.OrderOption
	inters         []Interceptor
	predicates     []predicate.StorageProvider
	withCloudfiles *CloudFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StorageProviderQuery builder.
func (spq *StorageProviderQuery) Where(ps ...predicate.StorageProvider) *StorageProviderQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *StorageProviderQuery) Limit(limit int) *StorageProviderQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *StorageProviderQuery) Offset(offset int) *StorageProviderQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *StorageProviderQuery) Unique(unique bool) *StorageProviderQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *StorageProviderQuery) Order(o ...storageprovider.OrderOption) *StorageProviderQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryCloudfiles chains the current query on the "cloudfiles" edge.
func (spq *StorageProviderQuery) QueryCloudfiles() *CloudFileQuery {
	query := (&CloudFileClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(storageprovider.Table, storageprovider.FieldID, selector),
			sqlgraph.To(cloudfile.Table, cloudfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, storageprovider.CloudfilesTable, storageprovider.CloudfilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StorageProvider entity from the query.
// Returns a *NotFoundError when no StorageProvider was found.
func (spq *StorageProviderQuery) First(ctx context.Context) (*StorageProvider, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{storageprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *StorageProviderQuery) FirstX(ctx context.Context) *StorageProvider {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StorageProvider ID from the query.
// Returns a *NotFoundError when no StorageProvider ID was found.
func (spq *StorageProviderQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{storageprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *StorageProviderQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StorageProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StorageProvider entity is found.
// Returns a *NotFoundError when no StorageProvider entities are found.
func (spq *StorageProviderQuery) Only(ctx context.Context) (*StorageProvider, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{storageprovider.Label}
	default:
		return nil, &NotSingularError{storageprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *StorageProviderQuery) OnlyX(ctx context.Context) *StorageProvider {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StorageProvider ID in the query.
// Returns a *NotSingularError when more than one StorageProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *StorageProviderQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{storageprovider.Label}
	default:
		err = &NotSingularError{storageprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *StorageProviderQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StorageProviders.
func (spq *StorageProviderQuery) All(ctx context.Context) ([]*StorageProvider, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StorageProvider, *StorageProviderQuery]()
	return withInterceptors[[]*StorageProvider](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *StorageProviderQuery) AllX(ctx context.Context) []*StorageProvider {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StorageProvider IDs.
func (spq *StorageProviderQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(storageprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *StorageProviderQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *StorageProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*StorageProviderQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *StorageProviderQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *StorageProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *StorageProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StorageProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *StorageProviderQuery) Clone() *StorageProviderQuery {
	if spq == nil {
		return nil
	}
	return &StorageProviderQuery{
		config:         spq.config,
		ctx:            spq.ctx.Clone(),
		order:          append([]storageprovider.OrderOption{}, spq.order...),
		inters:         append([]Interceptor{}, spq.inters...),
		predicates:     append([]predicate.StorageProvider{}, spq.predicates...),
		withCloudfiles: spq.withCloudfiles.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithCloudfiles tells the query-builder to eager-load the nodes that are connected to
// the "cloudfiles" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StorageProviderQuery) WithCloudfiles(opts ...func(*CloudFileQuery)) *StorageProviderQuery {
	query := (&CloudFileClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withCloudfiles = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StorageProvider.Query().
//		GroupBy(storageprovider.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *StorageProviderQuery) GroupBy(field string, fields ...string) *StorageProviderGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StorageProviderGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = storageprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.StorageProvider.Query().
//		Select(storageprovider.FieldCreatedAt).
//		Scan(ctx, &v)
func (spq *StorageProviderQuery) Select(fields ...string) *StorageProviderSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &StorageProviderSelect{StorageProviderQuery: spq}
	sbuild.label = storageprovider.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StorageProviderSelect configured with the given aggregations.
func (spq *StorageProviderQuery) Aggregate(fns ...AggregateFunc) *StorageProviderSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *StorageProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !storageprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *StorageProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StorageProvider, error) {
	var (
		nodes       = []*StorageProvider{}
		_spec       = spq.querySpec()
		loadedTypes = [1]bool{
			spq.withCloudfiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StorageProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StorageProvider{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withCloudfiles; query != nil {
		if err := spq.loadCloudfiles(ctx, query, nodes,
			func(n *StorageProvider) { n.Edges.Cloudfiles = []*CloudFile{} },
			func(n *StorageProvider, e *CloudFile) { n.Edges.Cloudfiles = append(n.Edges.Cloudfiles, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *StorageProviderQuery) loadCloudfiles(ctx context.Context, query *CloudFileQuery, nodes []*StorageProvider, init func(*StorageProvider), assign func(*StorageProvider, *CloudFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*StorageProvider)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CloudFile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(storageprovider.CloudfilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.cloud_file_storage_providers
		if fk == nil {
			return fmt.Errorf(`foreign-key "cloud_file_storage_providers" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "cloud_file_storage_providers" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (spq *StorageProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *StorageProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(storageprovider.Table, storageprovider.Columns, sqlgraph.NewFieldSpec(storageprovider.FieldID, field.TypeUint64))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storageprovider.FieldID)
		for i := range fields {
			if fields[i] != storageprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *StorageProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(storageprovider.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = storageprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StorageProviderGroupBy is the group-by builder for StorageProvider entities.
type StorageProviderGroupBy struct {
	selector
	build *StorageProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *StorageProviderGroupBy) Aggregate(fns ...AggregateFunc) *StorageProviderGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *StorageProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StorageProviderQuery, *StorageProviderGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *StorageProviderGroupBy) sqlScan(ctx context.Context, root *StorageProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StorageProviderSelect is the builder for selecting fields of StorageProvider entities.
type StorageProviderSelect struct {
	*StorageProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *StorageProviderSelect) Aggregate(fns ...AggregateFunc) *StorageProviderSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *StorageProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StorageProviderQuery, *StorageProviderSelect](ctx, sps.StorageProviderQuery, sps, sps.inters, v)
}

func (sps *StorageProviderSelect) sqlScan(ctx context.Context, root *StorageProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
