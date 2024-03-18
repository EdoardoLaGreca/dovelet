// Package credentials provides credential retrieval and management
package credentials

import (
	"sync"

	"google.golang.org/api/option"
)

// A Provider is any component which provides credentials.
type CredentialsProvider interface {
	// Provide returns the ClientOption associated with the credentials.
	// It returns an error if the credentials could not be fetched.
	Provide() (option.ClientOption, error)
}

// A Credentials provides synchronous safe retrieval of service account
// credentials Value.
type Credentials struct {
	m        sync.Mutex
	provider CredentialsProvider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider CredentialsProvider) *Credentials {
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
