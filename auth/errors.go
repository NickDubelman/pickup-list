package auth

// NotAuthorized is an error for when a user is not authorized to do something
type NotAuthorized struct{}

func (e NotAuthorized) Error() string {
	return "not authorized"
}

// TokenExpired is an error for when a user's access token is expired
type TokenExpired struct{}

func (e TokenExpired) Error() string {
	return "access token is expired"
}
