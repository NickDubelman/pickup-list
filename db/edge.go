// Code generated by entc, DO NOT EDIT.

package db

import "context"

func (l *List) Owner(ctx context.Context) (*User, error) {
	result, err := l.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryOwner().Only(ctx)
	}
	return result, err
}

func (l *List) Users(ctx context.Context) ([]*User, error) {
	result, err := l.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryUsers().All(ctx)
	}
	return result, err
}

func (np *NBAPlayer) User(ctx context.Context) (*User, error) {
	result, err := np.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = np.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) NbaPlayer(ctx context.Context) (*NBAPlayer, error) {
	result, err := u.Edges.NbaPlayerOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryNbaPlayer().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) OwnedLists(ctx context.Context) ([]*List, error) {
	result, err := u.Edges.OwnedListsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryOwnedLists().All(ctx)
	}
	return result, err
}

func (u *User) Lists(ctx context.Context) ([]*List, error) {
	result, err := u.Edges.ListsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryLists().All(ctx)
	}
	return result, err
}
