package graph

import "github.com/dom-m17/lms/backend/internal/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Querier db.Querier
}
