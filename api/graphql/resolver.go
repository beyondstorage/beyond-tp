//go:generate go run github.com/99designs/gqlgen ./api/graphql/gqlgen.yml

package graphql

import (
	"github.com/aos-dev/noah/task"

	"github.com/aos-dev/dm/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB     *models.DB
	Portal *task.Portal
}
