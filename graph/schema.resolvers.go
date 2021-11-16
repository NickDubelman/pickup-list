package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NickDubelman/pickup-list/auth"
	"github.com/NickDubelman/pickup-list/db"
	"github.com/NickDubelman/pickup-list/db/nbaplayer"
	"github.com/NickDubelman/pickup-list/db/user"
	"github.com/NickDubelman/pickup-list/graph/generated"
	"github.com/NickDubelman/pickup-list/graph/model"
)

func (r *mutationResolver) CreateList(ctx context.Context, input model.CreateListInput) (*db.List, error) {
	owner, err := r.client.User.Query().
		Where(user.Email(testUser)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return r.client.List.Create().
		SetName(input.Name).
		SetOwner(owner).
		Save(ctx)
}

func (r *mutationResolver) JoinList(ctx context.Context, input model.JoinListInput) (*db.List, error) {
	user, err := r.client.User.Query().
		Where(user.Email(testUser)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return r.client.List.UpdateOneID(input.ListID).AddUsers(user).Save(ctx)
}

func (r *mutationResolver) UnjoinList(ctx context.Context, input model.JoinListInput) (*db.List, error) {
	user, err := r.client.User.Query().
		Where(user.Email(testUser)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return r.client.List.UpdateOneID(input.ListID).RemoveUsers(user).Save(ctx)
}

func (r *mutationResolver) SetUser(ctx context.Context, input model.SetUserInput) (*db.User, error) {
	user, err := r.client.User.Query().
		Where(user.Email(testUser)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	update := user.Update().SetRealName(input.RealName).ClearNbaPlayer()

	if input.NbaName != "" {
		nbaPlayer, err := r.client.NBAPlayer.Query().
			Where(nbaplayer.Name((input.NbaName))).
			WithUser().
			Only(ctx)

		if err != nil {
			return nil, err
		}

		if nbaPlayer.Edges.User != nil {
			suffix := "Choose a player who is not already represented"
			user := nbaPlayer.Edges.User.RealName
			return nil, fmt.Errorf(
				"%s is already represented by %s. %s", input.NbaName, user, suffix,
			)
		}

		update = update.SetNbaPlayer(nbaPlayer)
	}

	return update.Save(ctx)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	return auth.RefreshAccessToken(ctx, input.RefreshToken)
}

func (r *queryResolver) User(ctx context.Context) (*db.User, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.User.Get(ctx, user.ID())
}

func (r *queryResolver) Lists(ctx context.Context) ([]*db.List, error) {
	return r.client.List.Query().All(ctx)
}

func (r *queryResolver) NbaPlayers(ctx context.Context) ([]*db.NBAPlayer, error) {
	return r.client.NBAPlayer.Query().All(ctx)
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
