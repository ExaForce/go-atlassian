// Package models provides the data structures used in the user token management.
package models

// UserTokensScheme represents a user token.
// ID is the unique identifier of the token.
// Label is the label assigned to the token.
// CreatedAt is the time when the token was created.
// LastAccess is the last time the token was accessed.
type UserTokensScheme struct {
	ID             string `json:"id,omitempty"`             // The ID of the user token.
	Label          string `json:"label,omitempty"`          // The label of the user token.
	CreatedAt      string `json:"createdAt,omitempty"`      // The creation time of the user token.
	LastAccess     string `json:"lastAccess,omitempty"`     // The last access time of the user token.
	DisabledStatus bool   `json:"disabledStatus,omitempty"` // The disabled status of the user token.
	Expiry         string `json:"expiry,omitempty"`         // The expiry time of the user token.
}

// OrgKeyPageScheme represents a page of organization API keys.
type OrgKeyPageScheme struct {
	Data  []*OrgKeyScheme      `json:"data,omitempty"`  // The keys on this page.
	Links *LinkPageModelScheme `json:"links,omitempty"` // Links to other pages.
}

// OrgKeyScheme represents an organization API Key.
type OrgKeyScheme struct {
	ID             string            `json:"id,omitempty"`             // The ID of the org key token.
	Label          string            `json:"label,omitempty"`          // The label of the org key token.
	CreatedAt      string            `json:"createdAt,omitempty"`      // The creation time of the org key token.
	LastActiveAt   string            `json:"lastActiveAt,omitempty"`   // The last active time of the org key token.
	ExpiresAt      string            `json:"expiresAt,omitempty"`      // The expiry time of the org key token.
	Token          *string           `json:"token,omitempty"`          // The token value (may be null).
	Resource       string            `json:"resource,omitempty"`       // The resource ARI associated with the token.
	OrgID          string            `json:"orgId,omitempty"`          // The organization ID.
	CreatedBy      *OrgKeyUserScheme `json:"createdBy,omitempty"`      // The user who created the token.
	Scopes         []string          `json:"scopes,omitempty"`         // The scopes granted to the token.
	RevokedBy      *OrgKeyUserScheme `json:"revokedBy,omitempty"`      // The user who revoked the token (may be null).
	RevocationDate *string           `json:"revocationDate,omitempty"` // The date the token was revoked (may be null).
}

// OrgKeyUserScheme represents a user associated with an org api-key (creator or revoker).
type OrgKeyUserScheme struct {
	ID         string  `json:"id,omitempty"`         // The ID of the user.
	Name       string  `json:"name,omitempty"`       // The name of the user.
	Email      string  `json:"email,omitempty"`      // The email of the user.
	AppType    *string `json:"appType,omitempty"`    // The app type (may be null).
	OrgID      *string `json:"orgId,omitempty"`      // The organization ID (may be null).
	UserStatus string  `json:"userStatus,omitempty"` // The status of the user.
}

// OrgTokenPageScheme represents a page of organization API tokens.
type OrgTokenPageScheme struct {
	Data  []*OrgTokenScheme    `json:"data,omitempty"`  // The tokens on this page.
	Links *LinkPageModelScheme `json:"links,omitempty"` // Links to other pages.
}

// OrgTokenScheme represents an organization API Token.
type OrgTokenScheme struct {
	ID           string            `json:"id,omitempty"`           // The ID of the org token.
	Label        string            `json:"label,omitempty"`        // The label of the org token.
	CreatedAt    string            `json:"createdAt,omitempty"`    // The creation time of the org token.
	LastActiveAt string            `json:"lastActiveAt,omitempty"` // The last active time of the org token.
	ExpiresAt    string            `json:"expiresAt,omitempty"`    // The expiry time of the org token.
	Status       string            `json:"status,omitempty"`       // The status of the org token (e.g., ALLOWED).
	Scopes       []string          `json:"scopes,omitempty"`       // The scopes granted to the token.
	User         *OrgKeyUserScheme `json:"user,omitempty"`         // The user associated with the token.
	CreatedBy    *OrgKeyUserScheme `json:"createdBy,omitempty"`    // The user who created the token (may be null).
}
