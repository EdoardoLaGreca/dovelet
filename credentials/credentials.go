// Package credentials provides credential retrieval and management
package credentials

import (
	"sync"

	"google.golang.org/api/option"
)

// A Provider is the interface for any component which will provide credentials
// Value.
type Provider interface {
	// Provide returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Provide() (option.ClientOption, error)
}

// A Credentials provides synchronous safe retrieval of service account
// credentials Value.
type Credentials struct {
	m        sync.Mutex
	provider Provider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider Provider) *Credentials {
	return &Credentials{
		provider: provider,
	}
}

// Get returns the credentials value, or error if the credentials Value failed
// to be retrieved.
func (c *Credentials) Get() (option.ClientOption, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.provider.Provide()
}
