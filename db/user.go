// Code generated by entc, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/NickDubelman/pickup-list/db/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RealName holds the value of the "real_name" field.
	RealName string `json:"real_name,omitempty"`
	// NbaName holds the value of the "nba_name" field.
	NbaName string `json:"nba_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// OwnedLists holds the value of the owned_lists edge.
	OwnedLists []*List `json:"owned_lists,omitempty"`
	// Lists holds the value of the lists edge.
	Lists []*List `json:"lists,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnedListsOrErr returns the OwnedLists value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnedListsOrErr() ([]*List, error) {
	if e.loadedTypes[0] {
		return e.OwnedLists, nil
	}
	return nil, &NotLoadedError{edge: "owned_lists"}
}

// ListsOrErr returns the Lists value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ListsOrErr() ([]*List, error) {
	if e.loadedTypes[1] {
		return e.Lists, nil
	}
	return nil, &NotLoadedError{edge: "lists"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldRealName, user.FieldNbaName, user.FieldEmail:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				u.RealName = value.String
			}
		case user.FieldNbaName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nba_name", values[i])
			} else if value.Valid {
				u.NbaName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOwnedLists queries the "owned_lists" edge of the User entity.
func (u *User) QueryOwnedLists() *ListQuery {
	return (&UserClient{config: u.config}).QueryOwnedLists(u)
}

// QueryLists queries the "lists" edge of the User entity.
func (u *User) QueryLists() *ListQuery {
	return (&UserClient{config: u.config}).QueryLists(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("db: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", real_name=")
	builder.WriteString(u.RealName)
	builder.WriteString(", nba_name=")
	builder.WriteString(u.NbaName)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}