package resolver

import (
	"context"

	"github.com/TakakiAraki09/k8s-lesson/internal/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Ctx     context.Context
	Queries *database.Queries
}
