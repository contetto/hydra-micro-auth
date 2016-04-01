package hydraauth

import (
	"github.com/micro/go-platform/auth"
	"golang.org/x/net/context"
)

// This package is implementing following interface of micro-platform/auth:
/*
// Auth handles client side validation of authentication
// The client does not actually handle authentication itself.
// This could be an oauth2 provider, openid, basic auth, etc.
type Auth interface {
	// Determine if a request with context is authorised
	// Should extract token from the context, check with
	// the authorizer and return an err if not authed.
	// Can be used for both client and server
	Authorized(ctx context.Context, req Request) (*Token, error)
	// Retrieve a token for this client, should handle refreshing
	Token() (*Token, error)
	// Lookup a token
	Introspect(ctx context.Context) (*Token, error)
	// Revoke a token
	Revoke(t *Token) error
	// Will retrieve token from the context
	FromContext(ctx context.Context) (*Token, bool)
	// Creates a context with the token which can be
	NewContext(ctx context.Context, t *Token) context.Context
	// Retrieves token from headers
	// We may get back a partial token here
	FromHeader(map[string]string) (*Token, bool)
	// Adds token to headers
	NewHeader(map[string]string, *Token) map[string]string
	// We cache policies locally from the auth server
	Start() error
	Stop() error
	// Name
	String() string
}
*/

type Auth struct {
	BaseURL string
}

func New(baseURL string) Auth {
	return Auth{BaseURL: baseURL}
}

// Determine if a request with context is authorised
// Should extract token from the context, check with
// the authorizer and return an err if not authed.
// Can be used for both client and server
func (a Auth) Authorized(ctx context.Context, req auth.Request) (*auth.Token, error) {
	return &auth.Token{}, nil
}

// Retrieve a token for this client, should handle refreshing
func (a Auth) Token() (*auth.Token, error) {
	return &auth.Token{}, nil
}

// Lookup a token
func (a Auth) Introspect(ctx context.Context) (*auth.Token, error) {
	return &auth.Token{}, nil
}

// Revoke a token
func (a Auth) Revoke(t *auth.Token) error {
	return nil
}

// Will retrieve token from the context
func (a Auth) FromContext(ctx context.Context) (*auth.Token, bool) {
	return &auth.Token{}, true
}

// Creates a context with the token which can be
func (a Auth) NewContext(ctx context.Context, t *auth.Token) context.Context {
	return ctx
}

// Retrieves token from headers
// We may get back a partial token here{
func (a Auth) FromHeader(map[string]string) (*auth.Token, bool) {
	return &auth.Token{}, true
}

// Adds token to headers
func (a Auth) NewHeader(map[string]string, *auth.Token) map[string]string {
	stringMap := make(map[string]string)
	return stringMap
}

// We cache policies locally from the auth server
func (a Auth) Start() error {
	return nil
}

func (a Auth) Stop() error {
	return nil
}

// Name
func (a Auth) String() string {
	return ""
}
