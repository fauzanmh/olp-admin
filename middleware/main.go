package middleware

import (
	appInit "github.com/fauzanmh/olp-admin/init"
)

// Middleware for
type Middleware struct {
	config *appInit.Config
}

// NewMiddleware for
func NewMiddleware(config *appInit.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}
