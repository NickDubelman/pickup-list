package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/NickDubelman/pickup-list/db"
	"github.com/NickDubelman/pickup-list/graph/generated"
	"github.com/NickDubelman/pickup-list/graph/model"
)

func (r *mutationResolver) CreateList(ctx context.Context, input model.CreateListInput) (*db.List, error) {
	ownerID := 4294967296

	return r.client.List.Create().
		SetName(input.Name).
		SetOwnerID(ownerID).
		Save(ctx)
}

func (r *mutationResolver) JoinList(ctx context.Context, input model.JoinListInput) (*db.List, error) {
	userID := 4294967296

	return r.client.List.UpdateOneID(input.ListID).AddUserIDs(userID).Save(ctx)
}

func (r *queryResolver) Lists(ctx context.Context) ([]*db.List, error) {
	return r.client.List.Query().All(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (db.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]db.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
