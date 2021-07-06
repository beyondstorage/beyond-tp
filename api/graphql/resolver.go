package graphql

import (
	"github.com/beyondstorage/beyond-tp/models"
	"github.com/beyondstorage/beyond-tp/task"
)

type Resolver struct {
	DB      *models.DB
	Manager *task.Manager
}
