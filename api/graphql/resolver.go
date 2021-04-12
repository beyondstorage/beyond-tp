package graphql

import (
	"github.com/aos-dev/dm/models"
	"github.com/aos-dev/dm/task"
)

type Resolver struct {
	DB      *models.DB
	Manager *task.Manager
}
