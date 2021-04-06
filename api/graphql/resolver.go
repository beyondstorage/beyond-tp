package graphql

import (
	"github.com/aos-dev/dm/task"

	"github.com/aos-dev/dm/models"
)

type Resolver struct {
	DB      *models.DB
	Manager *task.Manager
}
