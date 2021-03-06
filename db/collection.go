// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *ListQuery) CollectFields(ctx context.Context, satisfies ...string) *ListQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		l = l.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return l
}

func (l *ListQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ListQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "owner":
			l = l.WithOwner(func(query *UserQuery) {
				query.collectField(ctx, field)
			})
		case "users":
			l = l.WithUsers(func(query *UserQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return l
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (np *NBAPlayerQuery) CollectFields(ctx context.Context, satisfies ...string) *NBAPlayerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		np = np.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return np
}

func (np *NBAPlayerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *NBAPlayerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			np = np.WithUser(func(query *UserQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return np
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "lists":
			u = u.WithLists(func(query *ListQuery) {
				query.collectField(ctx, field)
			})
		case "nba_player":
			u = u.WithNbaPlayer(func(query *NBAPlayerQuery) {
				query.collectField(ctx, field)
			})
		case "owned_lists":
			u = u.WithOwnedLists(func(query *ListQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return u
}
