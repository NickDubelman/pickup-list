// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NickDubelman/pickup-list/db/nbaplayer"
	"github.com/NickDubelman/pickup-list/db/predicate"
	"github.com/NickDubelman/pickup-list/db/user"
)

// NBAPlayerQuery is the builder for querying NBAPlayer entities.
type NBAPlayerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NBAPlayer
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NBAPlayerQuery builder.
func (npq *NBAPlayerQuery) Where(ps ...predicate.NBAPlayer) *NBAPlayerQuery {
	npq.predicates = append(npq.predicates, ps...)
	return npq
}

// Limit adds a limit step to the query.
func (npq *NBAPlayerQuery) Limit(limit int) *NBAPlayerQuery {
	npq.limit = &limit
	return npq
}

// Offset adds an offset step to the query.
func (npq *NBAPlayerQuery) Offset(offset int) *NBAPlayerQuery {
	npq.offset = &offset
	return npq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (npq *NBAPlayerQuery) Unique(unique bool) *NBAPlayerQuery {
	npq.unique = &unique
	return npq
}

// Order adds an order step to the query.
func (npq *NBAPlayerQuery) Order(o ...OrderFunc) *NBAPlayerQuery {
	npq.order = append(npq.order, o...)
	return npq
}

// QueryUser chains the current query on the "user" edge.
func (npq *NBAPlayerQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: npq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := npq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nbaplayer.Table, nbaplayer.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, nbaplayer.UserTable, nbaplayer.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(npq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NBAPlayer entity from the query.
// Returns a *NotFoundError when no NBAPlayer was found.
func (npq *NBAPlayerQuery) First(ctx context.Context) (*NBAPlayer, error) {
	nodes, err := npq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nbaplayer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (npq *NBAPlayerQuery) FirstX(ctx context.Context) *NBAPlayer {
	node, err := npq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NBAPlayer ID from the query.
// Returns a *NotFoundError when no NBAPlayer ID was found.
func (npq *NBAPlayerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = npq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nbaplayer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (npq *NBAPlayerQuery) FirstIDX(ctx context.Context) int {
	id, err := npq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NBAPlayer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NBAPlayer entity is not found.
// Returns a *NotFoundError when no NBAPlayer entities are found.
func (npq *NBAPlayerQuery) Only(ctx context.Context) (*NBAPlayer, error) {
	nodes, err := npq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nbaplayer.Label}
	default:
		return nil, &NotSingularError{nbaplayer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (npq *NBAPlayerQuery) OnlyX(ctx context.Context) *NBAPlayer {
	node, err := npq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NBAPlayer ID in the query.
// Returns a *NotSingularError when exactly one NBAPlayer ID is not found.
// Returns a *NotFoundError when no entities are found.
func (npq *NBAPlayerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = npq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = &NotSingularError{nbaplayer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (npq *NBAPlayerQuery) OnlyIDX(ctx context.Context) int {
	id, err := npq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NBAPlayers.
func (npq *NBAPlayerQuery) All(ctx context.Context) ([]*NBAPlayer, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return npq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (npq *NBAPlayerQuery) AllX(ctx context.Context) []*NBAPlayer {
	nodes, err := npq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NBAPlayer IDs.
func (npq *NBAPlayerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := npq.Select(nbaplayer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (npq *NBAPlayerQuery) IDsX(ctx context.Context) []int {
	ids, err := npq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (npq *NBAPlayerQuery) Count(ctx context.Context) (int, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return npq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (npq *NBAPlayerQuery) CountX(ctx context.Context) int {
	count, err := npq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (npq *NBAPlayerQuery) Exist(ctx context.Context) (bool, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return npq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (npq *NBAPlayerQuery) ExistX(ctx context.Context) bool {
	exist, err := npq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NBAPlayerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (npq *NBAPlayerQuery) Clone() *NBAPlayerQuery {
	if npq == nil {
		return nil
	}
	return &NBAPlayerQuery{
		config:     npq.config,
		limit:      npq.limit,
		offset:     npq.offset,
		order:      append([]OrderFunc{}, npq.order...),
		predicates: append([]predicate.NBAPlayer{}, npq.predicates...),
		withUser:   npq.withUser.Clone(),
		// clone intermediate query.
		sql:  npq.sql.Clone(),
		path: npq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (npq *NBAPlayerQuery) WithUser(opts ...func(*UserQuery)) *NBAPlayerQuery {
	query := &UserQuery{config: npq.config}
	for _, opt := range opts {
		opt(query)
	}
	npq.withUser = query
	return npq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NBAPlayer.Query().
//		GroupBy(nbaplayer.FieldName).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (npq *NBAPlayerQuery) GroupBy(field string, fields ...string) *NBAPlayerGroupBy {
	group := &NBAPlayerGroupBy{config: npq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return npq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.NBAPlayer.Query().
//		Select(nbaplayer.FieldName).
//		Scan(ctx, &v)
//
func (npq *NBAPlayerQuery) Select(fields ...string) *NBAPlayerSelect {
	npq.fields = append(npq.fields, fields...)
	return &NBAPlayerSelect{NBAPlayerQuery: npq}
}

func (npq *NBAPlayerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range npq.fields {
		if !nbaplayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if npq.path != nil {
		prev, err := npq.path(ctx)
		if err != nil {
			return err
		}
		npq.sql = prev
	}
	return nil
}

func (npq *NBAPlayerQuery) sqlAll(ctx context.Context) ([]*NBAPlayer, error) {
	var (
		nodes       = []*NBAPlayer{}
		withFKs     = npq.withFKs
		_spec       = npq.querySpec()
		loadedTypes = [1]bool{
			npq.withUser != nil,
		}
	)
	if npq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, nbaplayer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NBAPlayer{config: npq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("db: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, npq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := npq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*NBAPlayer)
		for i := range nodes {
			if nodes[i].user_nba_player == nil {
				continue
			}
			fk := *nodes[i].user_nba_player
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_nba_player" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (npq *NBAPlayerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := npq.querySpec()
	return sqlgraph.CountNodes(ctx, npq.driver, _spec)
}

func (npq *NBAPlayerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := npq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (npq *NBAPlayerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nbaplayer.Table,
			Columns: nbaplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nbaplayer.FieldID,
			},
		},
		From:   npq.sql,
		Unique: true,
	}
	if unique := npq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := npq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nbaplayer.FieldID)
		for i := range fields {
			if fields[i] != nbaplayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := npq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := npq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := npq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := npq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (npq *NBAPlayerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(npq.driver.Dialect())
	t1 := builder.Table(nbaplayer.Table)
	columns := npq.fields
	if len(columns) == 0 {
		columns = nbaplayer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if npq.sql != nil {
		selector = npq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range npq.predicates {
		p(selector)
	}
	for _, p := range npq.order {
		p(selector)
	}
	if offset := npq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := npq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NBAPlayerGroupBy is the group-by builder for NBAPlayer entities.
type NBAPlayerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (npgb *NBAPlayerGroupBy) Aggregate(fns ...AggregateFunc) *NBAPlayerGroupBy {
	npgb.fns = append(npgb.fns, fns...)
	return npgb
}

// Scan applies the group-by query and scans the result into the given value.
func (npgb *NBAPlayerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := npgb.path(ctx)
	if err != nil {
		return err
	}
	npgb.sql = query
	return npgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := npgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(npgb.fields) > 1 {
		return nil, errors.New("db: NBAPlayerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := npgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) StringsX(ctx context.Context) []string {
	v, err := npgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = npgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) StringX(ctx context.Context) string {
	v, err := npgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(npgb.fields) > 1 {
		return nil, errors.New("db: NBAPlayerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := npgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) IntsX(ctx context.Context) []int {
	v, err := npgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = npgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) IntX(ctx context.Context) int {
	v, err := npgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(npgb.fields) > 1 {
		return nil, errors.New("db: NBAPlayerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := npgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := npgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = npgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := npgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(npgb.fields) > 1 {
		return nil, errors.New("db: NBAPlayerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := npgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := npgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (npgb *NBAPlayerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = npgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (npgb *NBAPlayerGroupBy) BoolX(ctx context.Context) bool {
	v, err := npgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (npgb *NBAPlayerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range npgb.fields {
		if !nbaplayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := npgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := npgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (npgb *NBAPlayerGroupBy) sqlQuery() *sql.Selector {
	selector := npgb.sql.Select()
	aggregation := make([]string, 0, len(npgb.fns))
	for _, fn := range npgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(npgb.fields)+len(npgb.fns))
		for _, f := range npgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(npgb.fields...)...)
}

// NBAPlayerSelect is the builder for selecting fields of NBAPlayer entities.
type NBAPlayerSelect struct {
	*NBAPlayerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nps *NBAPlayerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nps.prepareQuery(ctx); err != nil {
		return err
	}
	nps.sql = nps.NBAPlayerQuery.sqlQuery(ctx)
	return nps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nps *NBAPlayerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nps.fields) > 1 {
		return nil, errors.New("db: NBAPlayerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nps *NBAPlayerSelect) StringsX(ctx context.Context) []string {
	v, err := nps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nps *NBAPlayerSelect) StringX(ctx context.Context) string {
	v, err := nps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nps.fields) > 1 {
		return nil, errors.New("db: NBAPlayerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nps *NBAPlayerSelect) IntsX(ctx context.Context) []int {
	v, err := nps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nps *NBAPlayerSelect) IntX(ctx context.Context) int {
	v, err := nps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nps.fields) > 1 {
		return nil, errors.New("db: NBAPlayerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nps *NBAPlayerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nps *NBAPlayerSelect) Float64X(ctx context.Context) float64 {
	v, err := nps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nps.fields) > 1 {
		return nil, errors.New("db: NBAPlayerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nps *NBAPlayerSelect) BoolsX(ctx context.Context) []bool {
	v, err := nps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nps *NBAPlayerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nbaplayer.Label}
	default:
		err = fmt.Errorf("db: NBAPlayerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nps *NBAPlayerSelect) BoolX(ctx context.Context) bool {
	v, err := nps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nps *NBAPlayerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nps.sql.Query()
	if err := nps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
