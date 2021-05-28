package graphql

import (
	"github.com/beyondstorage/dm/models"
	"github.com/beyondstorage/dm/task"
)

type Resolver struct {
	DB      *models.DB
	Manager *task.Manager
}
