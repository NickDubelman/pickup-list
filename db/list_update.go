// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NickDubelman/pickup-list/db/list"
	"github.com/NickDubelman/pickup-list/db/predicate"
	"github.com/NickDubelman/pickup-list/db/user"
)

// ListUpdate is the builder for updating List entities.
type ListUpdate struct {
	config
	hooks    []Hook
	mutation *ListMutation
}

// Where appends a list predicates to the ListUpdate builder.
func (lu *ListUpdate) Where(ps ...predicate.List) *ListUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *ListUpdate) SetName(s string) *ListUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *ListUpdate) SetCreatedAt(t time.Time) *ListUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *ListUpdate) SetNillableCreatedAt(t *time.Time) *ListUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lu *ListUpdate) SetOwnerID(id int) *ListUpdate {
	lu.mutation.SetOwnerID(id)
	return lu
}

// SetOwner sets the "owner" edge to the User entity.
func (lu *ListUpdate) SetOwner(u *User) *ListUpdate {
	return lu.SetOwnerID(u.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (lu *ListUpdate) AddUserIDs(ids ...int) *ListUpdate {
	lu.mutation.AddUserIDs(ids...)
	return lu
}

// AddUsers adds the "users" edges to the User entity.
func (lu *ListUpdate) AddUsers(u ...*User) *ListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.AddUserIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (lu *ListUpdate) Mutation() *ListMutation {
	return lu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (lu *ListUpdate) ClearOwner() *ListUpdate {
	lu.mutation.ClearOwner()
	return lu
}

// ClearUsers clears all "users" edges to the User entity.
func (lu *ListUpdate) ClearUsers() *ListUpdate {
	lu.mutation.ClearUsers()
	return lu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (lu *ListUpdate) RemoveUserIDs(ids ...int) *ListUpdate {
	lu.mutation.RemoveUserIDs(ids...)
	return lu
}

// RemoveUsers removes "users" edges to User entities.
func (lu *ListUpdate) RemoveUsers(u ...*User) *ListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *ListUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *ListUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *ListUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *ListUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *ListUpdate) check() error {
	if _, ok := lu.mutation.OwnerID(); lu.mutation.OwnerCleared() && !ok {
		return errors.New("db: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (lu *ListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: list.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: list.FieldName,
		})
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldCreatedAt,
		})
	}
	if lu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   list.OwnerTable,
			Columns: []string{list.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   list.OwnerTable,
			Columns: []string{list.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !lu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ListUpdateOne is the builder for updating a single List entity.
type ListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ListMutation
}

// SetName sets the "name" field.
func (luo *ListUpdateOne) SetName(s string) *ListUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *ListUpdateOne) SetCreatedAt(t time.Time) *ListUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *ListUpdateOne) SetNillableCreatedAt(t *time.Time) *ListUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (luo *ListUpdateOne) SetOwnerID(id int) *ListUpdateOne {
	luo.mutation.SetOwnerID(id)
	return luo
}

// SetOwner sets the "owner" edge to the User entity.
func (luo *ListUpdateOne) SetOwner(u *User) *ListUpdateOne {
	return luo.SetOwnerID(u.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (luo *ListUpdateOne) AddUserIDs(ids ...int) *ListUpdateOne {
	luo.mutation.AddUserIDs(ids...)
	return luo
}

// AddUsers adds the "users" edges to the User entity.
func (luo *ListUpdateOne) AddUsers(u ...*User) *ListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.AddUserIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (luo *ListUpdateOne) Mutation() *ListMutation {
	return luo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (luo *ListUpdateOne) ClearOwner() *ListUpdateOne {
	luo.mutation.ClearOwner()
	return luo
}

// ClearUsers clears all "users" edges to the User entity.
func (luo *ListUpdateOne) ClearUsers() *ListUpdateOne {
	luo.mutation.ClearUsers()
	return luo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (luo *ListUpdateOne) RemoveUserIDs(ids ...int) *ListUpdateOne {
	luo.mutation.RemoveUserIDs(ids...)
	return luo
}

// RemoveUsers removes "users" edges to User entities.
func (luo *ListUpdateOne) RemoveUsers(u ...*User) *ListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *ListUpdateOne) Select(field string, fields ...string) *ListUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated List entity.
func (luo *ListUpdateOne) Save(ctx context.Context) (*List, error) {
	var (
		err  error
		node *List
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *ListUpdateOne) SaveX(ctx context.Context) *List {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *ListUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *ListUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *ListUpdateOne) check() error {
	if _, ok := luo.mutation.OwnerID(); luo.mutation.OwnerCleared() && !ok {
		return errors.New("db: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (luo *ListUpdateOne) sqlSave(ctx context.Context) (_node *List, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: list.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing List.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, list.FieldID)
		for _, f := range fields {
			if !list.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != list.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: list.FieldName,
		})
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldCreatedAt,
		})
	}
	if luo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   list.OwnerTable,
			Columns: []string{list.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   list.OwnerTable,
			Columns: []string{list.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !luo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   list.UsersTable,
			Columns: list.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &List{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}