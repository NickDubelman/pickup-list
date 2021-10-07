// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NickDubelman/pickup-list/db/list"
	"github.com/NickDubelman/pickup-list/db/predicate"
	"github.com/NickDubelman/pickup-list/db/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeList = "List"
	TypeUser = "User"
)

// ListMutation represents an operation that mutates the List nodes in the graph.
type ListMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	users         map[int]struct{}
	removedusers  map[int]struct{}
	clearedusers  bool
	done          bool
	oldValue      func(context.Context) (*List, error)
	predicates    []predicate.List
}

var _ ent.Mutation = (*ListMutation)(nil)

// listOption allows management of the mutation configuration using functional options.
type listOption func(*ListMutation)

// newListMutation creates new mutation for the List entity.
func newListMutation(c config, op Op, opts ...listOption) *ListMutation {
	m := &ListMutation{
		config:        c,
		op:            op,
		typ:           TypeList,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withListID sets the ID field of the mutation.
func withListID(id int) listOption {
	return func(m *ListMutation) {
		var (
			err   error
			once  sync.Once
			value *List
		)
		m.oldValue = func(ctx context.Context) (*List, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().List.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withList sets the old List of the mutation.
func withList(node *List) listOption {
	return func(m *ListMutation) {
		m.oldValue = func(context.Context) (*List, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ListMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ListMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ListMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *ListMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ListMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the List entity.
// If the List object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ListMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
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

// ResetName resets all changes to the "name" field.
func (m *ListMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ListMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ListMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the List entity.
// If the List object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ListMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ListMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *ListMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *ListMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *ListMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *ListMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *ListMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *ListMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// AddUserIDs adds the "users" edge to the User entity by ids.
func (m *ListMutation) AddUserIDs(ids ...int) {
	if m.users == nil {
		m.users = make(map[int]struct{})
	}
	for i := range ids {
		m.users[ids[i]] = struct{}{}
	}
}

// ClearUsers clears the "users" edge to the User entity.
func (m *ListMutation) ClearUsers() {
	m.clearedusers = true
}

// UsersCleared reports if the "users" edge to the User entity was cleared.
func (m *ListMutation) UsersCleared() bool {
	return m.clearedusers
}

// RemoveUserIDs removes the "users" edge to the User entity by IDs.
func (m *ListMutation) RemoveUserIDs(ids ...int) {
	if m.removedusers == nil {
		m.removedusers = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.users, ids[i])
		m.removedusers[ids[i]] = struct{}{}
	}
}

// RemovedUsers returns the removed IDs of the "users" edge to the User entity.
func (m *ListMutation) RemovedUsersIDs() (ids []int) {
	for id := range m.removedusers {
		ids = append(ids, id)
	}
	return
}

// UsersIDs returns the "users" edge IDs in the mutation.
func (m *ListMutation) UsersIDs() (ids []int) {
	for id := range m.users {
		ids = append(ids, id)
	}
	return
}

// ResetUsers resets all changes to the "users" edge.
func (m *ListMutation) ResetUsers() {
	m.users = nil
	m.clearedusers = false
	m.removedusers = nil
}

// Where appends a list predicates to the ListMutation builder.
func (m *ListMutation) Where(ps ...predicate.List) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ListMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (List).
func (m *ListMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ListMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, list.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, list.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ListMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case list.FieldName:
		return m.Name()
	case list.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ListMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case list.FieldName:
		return m.OldName(ctx)
	case list.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown List field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ListMutation) SetField(name string, value ent.Value) error {
	switch name {
	case list.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case list.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown List field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ListMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ListMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ListMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown List numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ListMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ListMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ListMutation) ClearField(name string) error {
	return fmt.Errorf("unknown List nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ListMutation) ResetField(name string) error {
	switch name {
	case list.FieldName:
		m.ResetName()
		return nil
	case list.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown List field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ListMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.owner != nil {
		edges = append(edges, list.EdgeOwner)
	}
	if m.users != nil {
		edges = append(edges, list.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ListMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case list.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	case list.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.users))
		for id := range m.users {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ListMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedusers != nil {
		edges = append(edges, list.EdgeUsers)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ListMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case list.EdgeUsers:
		ids := make([]ent.Value, 0, len(m.removedusers))
		for id := range m.removedusers {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ListMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedowner {
		edges = append(edges, list.EdgeOwner)
	}
	if m.clearedusers {
		edges = append(edges, list.EdgeUsers)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ListMutation) EdgeCleared(name string) bool {
	switch name {
	case list.EdgeOwner:
		return m.clearedowner
	case list.EdgeUsers:
		return m.clearedusers
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ListMutation) ClearEdge(name string) error {
	switch name {
	case list.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown List unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ListMutation) ResetEdge(name string) error {
	switch name {
	case list.EdgeOwner:
		m.ResetOwner()
		return nil
	case list.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown List edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	real_name          *string
	nba_name           *string
	created_at         *time.Time
	clearedFields      map[string]struct{}
	owned_lists        map[int]struct{}
	removedowned_lists map[int]struct{}
	clearedowned_lists bool
	lists              map[int]struct{}
	removedlists       map[int]struct{}
	clearedlists       bool
	done               bool
	oldValue           func(context.Context) (*User, error)
	predicates         []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetRealName sets the "real_name" field.
func (m *UserMutation) SetRealName(s string) {
	m.real_name = &s
}

// RealName returns the value of the "real_name" field in the mutation.
func (m *UserMutation) RealName() (r string, exists bool) {
	v := m.real_name
	if v == nil {
		return
	}
	return *v, true
}

// OldRealName returns the old "real_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRealName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRealName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRealName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRealName: %w", err)
	}
	return oldValue.RealName, nil
}

// ResetRealName resets all changes to the "real_name" field.
func (m *UserMutation) ResetRealName() {
	m.real_name = nil
}

// SetNbaName sets the "nba_name" field.
func (m *UserMutation) SetNbaName(s string) {
	m.nba_name = &s
}

// NbaName returns the value of the "nba_name" field in the mutation.
func (m *UserMutation) NbaName() (r string, exists bool) {
	v := m.nba_name
	if v == nil {
		return
	}
	return *v, true
}

// OldNbaName returns the old "nba_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNbaName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNbaName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNbaName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNbaName: %w", err)
	}
	return oldValue.NbaName, nil
}

// ResetNbaName resets all changes to the "nba_name" field.
func (m *UserMutation) ResetNbaName() {
	m.nba_name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddOwnedListIDs adds the "owned_lists" edge to the List entity by ids.
func (m *UserMutation) AddOwnedListIDs(ids ...int) {
	if m.owned_lists == nil {
		m.owned_lists = make(map[int]struct{})
	}
	for i := range ids {
		m.owned_lists[ids[i]] = struct{}{}
	}
}

// ClearOwnedLists clears the "owned_lists" edge to the List entity.
func (m *UserMutation) ClearOwnedLists() {
	m.clearedowned_lists = true
}

// OwnedListsCleared reports if the "owned_lists" edge to the List entity was cleared.
func (m *UserMutation) OwnedListsCleared() bool {
	return m.clearedowned_lists
}

// RemoveOwnedListIDs removes the "owned_lists" edge to the List entity by IDs.
func (m *UserMutation) RemoveOwnedListIDs(ids ...int) {
	if m.removedowned_lists == nil {
		m.removedowned_lists = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.owned_lists, ids[i])
		m.removedowned_lists[ids[i]] = struct{}{}
	}
}

// RemovedOwnedLists returns the removed IDs of the "owned_lists" edge to the List entity.
func (m *UserMutation) RemovedOwnedListsIDs() (ids []int) {
	for id := range m.removedowned_lists {
		ids = append(ids, id)
	}
	return
}

// OwnedListsIDs returns the "owned_lists" edge IDs in the mutation.
func (m *UserMutation) OwnedListsIDs() (ids []int) {
	for id := range m.owned_lists {
		ids = append(ids, id)
	}
	return
}

// ResetOwnedLists resets all changes to the "owned_lists" edge.
func (m *UserMutation) ResetOwnedLists() {
	m.owned_lists = nil
	m.clearedowned_lists = false
	m.removedowned_lists = nil
}

// AddListIDs adds the "lists" edge to the List entity by ids.
func (m *UserMutation) AddListIDs(ids ...int) {
	if m.lists == nil {
		m.lists = make(map[int]struct{})
	}
	for i := range ids {
		m.lists[ids[i]] = struct{}{}
	}
}

// ClearLists clears the "lists" edge to the List entity.
func (m *UserMutation) ClearLists() {
	m.clearedlists = true
}

// ListsCleared reports if the "lists" edge to the List entity was cleared.
func (m *UserMutation) ListsCleared() bool {
	return m.clearedlists
}

// RemoveListIDs removes the "lists" edge to the List entity by IDs.
func (m *UserMutation) RemoveListIDs(ids ...int) {
	if m.removedlists == nil {
		m.removedlists = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.lists, ids[i])
		m.removedlists[ids[i]] = struct{}{}
	}
}

// RemovedLists returns the removed IDs of the "lists" edge to the List entity.
func (m *UserMutation) RemovedListsIDs() (ids []int) {
	for id := range m.removedlists {
		ids = append(ids, id)
	}
	return
}

// ListsIDs returns the "lists" edge IDs in the mutation.
func (m *UserMutation) ListsIDs() (ids []int) {
	for id := range m.lists {
		ids = append(ids, id)
	}
	return
}

// ResetLists resets all changes to the "lists" edge.
func (m *UserMutation) ResetLists() {
	m.lists = nil
	m.clearedlists = false
	m.removedlists = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.real_name != nil {
		fields = append(fields, user.FieldRealName)
	}
	if m.nba_name != nil {
		fields = append(fields, user.FieldNbaName)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldRealName:
		return m.RealName()
	case user.FieldNbaName:
		return m.NbaName()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldRealName:
		return m.OldRealName(ctx)
	case user.FieldNbaName:
		return m.OldNbaName(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldRealName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRealName(v)
		return nil
	case user.FieldNbaName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNbaName(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldRealName:
		m.ResetRealName()
		return nil
	case user.FieldNbaName:
		m.ResetNbaName()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.owned_lists != nil {
		edges = append(edges, user.EdgeOwnedLists)
	}
	if m.lists != nil {
		edges = append(edges, user.EdgeLists)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOwnedLists:
		ids := make([]ent.Value, 0, len(m.owned_lists))
		for id := range m.owned_lists {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeLists:
		ids := make([]ent.Value, 0, len(m.lists))
		for id := range m.lists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedowned_lists != nil {
		edges = append(edges, user.EdgeOwnedLists)
	}
	if m.removedlists != nil {
		edges = append(edges, user.EdgeLists)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeOwnedLists:
		ids := make([]ent.Value, 0, len(m.removedowned_lists))
		for id := range m.removedowned_lists {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeLists:
		ids := make([]ent.Value, 0, len(m.removedlists))
		for id := range m.removedlists {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedowned_lists {
		edges = append(edges, user.EdgeOwnedLists)
	}
	if m.clearedlists {
		edges = append(edges, user.EdgeLists)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeOwnedLists:
		return m.clearedowned_lists
	case user.EdgeLists:
		return m.clearedlists
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeOwnedLists:
		m.ResetOwnedLists()
		return nil
	case user.EdgeLists:
		m.ResetLists()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
