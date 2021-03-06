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
	"github.com/NickDubelman/pickup-list/db/user"
)

// ListCreate is the builder for creating a List entity.
type ListCreate struct {
	config
	mutation *ListMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (lc *ListCreate) SetName(s string) *ListCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *ListCreate) SetCreatedAt(t time.Time) *ListCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *ListCreate) SetNillableCreatedAt(t *time.Time) *ListCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (lc *ListCreate) SetOwnerID(id int) *ListCreate {
	lc.mutation.SetOwnerID(id)
	return lc
}

// SetOwner sets the "owner" edge to the User entity.
func (lc *ListCreate) SetOwner(u *User) *ListCreate {
	return lc.SetOwnerID(u.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (lc *ListCreate) AddUserIDs(ids ...int) *ListCreate {
	lc.mutation.AddUserIDs(ids...)
	return lc
}

// AddUsers adds the "users" edges to the User entity.
func (lc *ListCreate) AddUsers(u ...*User) *ListCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lc.AddUserIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (lc *ListCreate) Mutation() *ListMutation {
	return lc.mutation
}

// Save creates the List in the database.
func (lc *ListCreate) Save(ctx context.Context) (*List, error) {
	var (
		err  error
		node *List
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *ListCreate) SaveX(ctx context.Context) *List {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *ListCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *ListCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *ListCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := list.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *ListCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "name"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "created_at"`)}
	}
	if _, ok := lc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New("db: missing required edge \"owner\"")}
	}
	return nil
}

func (lc *ListCreate) sqlSave(ctx context.Context) (*List, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *ListCreate) createSpec() (*List, *sqlgraph.CreateSpec) {
	var (
		_node = &List{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: list.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: list.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: list.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: list.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := lc.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.list_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.List.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ListUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (lc *ListCreate) OnConflict(opts ...sql.ConflictOption) *ListUpsertOne {
	lc.conflict = opts
	return &ListUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.List.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *ListCreate) OnConflictColumns(columns ...string) *ListUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &ListUpsertOne{
		create: lc,
	}
}

type (
	// ListUpsertOne is the builder for "upsert"-ing
	//  one List node.
	ListUpsertOne struct {
		create *ListCreate
	}

	// ListUpsert is the "OnConflict" setter.
	ListUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ListUpsert) SetName(v string) *ListUpsert {
	u.Set(list.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ListUpsert) UpdateName() *ListUpsert {
	u.SetExcluded(list.FieldName)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ListUpsert) SetCreatedAt(v time.Time) *ListUpsert {
	u.Set(list.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ListUpsert) UpdateCreatedAt() *ListUpsert {
	u.SetExcluded(list.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.List.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ListUpsertOne) UpdateNewValues() *ListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.List.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ListUpsertOne) Ignore() *ListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ListUpsertOne) DoNothing() *ListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ListCreate.OnConflict
// documentation for more info.
func (u *ListUpsertOne) Update(set func(*ListUpsert)) *ListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ListUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ListUpsertOne) SetName(v string) *ListUpsertOne {
	return u.Update(func(s *ListUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ListUpsertOne) UpdateName() *ListUpsertOne {
	return u.Update(func(s *ListUpsert) {
		s.UpdateName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ListUpsertOne) SetCreatedAt(v time.Time) *ListUpsertOne {
	return u.Update(func(s *ListUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ListUpsertOne) UpdateCreatedAt() *ListUpsertOne {
	return u.Update(func(s *ListUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *ListUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for ListCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ListUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ListUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ListUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ListCreateBulk is the builder for creating many List entities in bulk.
type ListCreateBulk struct {
	config
	builders []*ListCreate
	conflict []sql.ConflictOption
}

// Save creates the List entities in the database.
func (lcb *ListCreateBulk) Save(ctx context.Context) ([]*List, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*List, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *ListCreateBulk) SaveX(ctx context.Context) []*List {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *ListCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *ListCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.List.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ListUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *ListCreateBulk) OnConflict(opts ...sql.ConflictOption) *ListUpsertBulk {
	lcb.conflict = opts
	return &ListUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.List.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *ListCreateBulk) OnConflictColumns(columns ...string) *ListUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &ListUpsertBulk{
		create: lcb,
	}
}

// ListUpsertBulk is the builder for "upsert"-ing
// a bulk of List nodes.
type ListUpsertBulk struct {
	create *ListCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.List.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ListUpsertBulk) UpdateNewValues() *ListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.List.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ListUpsertBulk) Ignore() *ListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ListUpsertBulk) DoNothing() *ListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ListCreateBulk.OnConflict
// documentation for more info.
func (u *ListUpsertBulk) Update(set func(*ListUpsert)) *ListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ListUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ListUpsertBulk) SetName(v string) *ListUpsertBulk {
	return u.Update(func(s *ListUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ListUpsertBulk) UpdateName() *ListUpsertBulk {
	return u.Update(func(s *ListUpsert) {
		s.UpdateName()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ListUpsertBulk) SetCreatedAt(v time.Time) *ListUpsertBulk {
	return u.Update(func(s *ListUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ListUpsertBulk) UpdateCreatedAt() *ListUpsertBulk {
	return u.Update(func(s *ListUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *ListUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the ListCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for ListCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ListUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
