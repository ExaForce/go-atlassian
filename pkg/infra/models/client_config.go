package models

import "time"

// ClientConfig holds configuration options for the Client
type ClientConfig struct {
	// MaxRetries is the maximum number of retries for rate limit handling
	MaxRetries int
	// InitialRetryDelay is the initial delay between retries in milliseconds
	InitialRetryDelay time.Duration
	// MaxRetryDelay is the maximum delay between retries in milliseconds
	MaxRetryDelay time.Duration
}
