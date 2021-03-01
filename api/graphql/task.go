package graphql

import (
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	"github.com/aos-dev/dm/model"
)

var taskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "task id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "task name",
		},
		"status": &graphql.Field{
			Type:        taskStatusEnum,
			Description: "task status",
		},
		"created_at": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "task create time",
		},
		"updated_at": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "task update time",
		},
	},
})

var taskStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Status",
	Description: "One of the status for task",
	Values: graphql.EnumValueConfigMap{
		"created": &graphql.EnumValueConfig{
			Value: model.StatusCreated,
		},
		"running": &graphql.EnumValueConfig{
			Value: model.StatusRunning,
		},
		"finished": &graphql.EnumValueConfig{
			Value: model.StatusFinished,
		},
		"stopped": &graphql.EnumValueConfig{
			Value: model.StatusStopped,
		},
		"broken": &graphql.EnumValueConfig{
			Value: model.StatusBroken,
		},
		"unknown": &graphql.EnumValueConfig{
			Value: model.StatusUnknown,
		},
	},
})

var getTask = graphql.Field{
	Type:        taskType,
	Description: "Get task by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "id of a task",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			// Find task
			task, err := model.GetTaskByID(id)
			if err != nil {
				return nil, err
			}
			return task, nil
		}
		return nil, nil
	},
}

var getTaskList = graphql.Field{
	Type:        graphql.NewList(taskType),
	Description: "Get task list",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		tasks, err := model.GetTaskList()
		if err != nil {
			return nil, err
		}
		return tasks, nil
	},
}

var createTask = graphql.Field{
	Type:        taskType,
	Description: "Create new task",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"status": &graphql.ArgumentConfig{
			Type: taskStatusEnum,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		now := time.Now()
		task := model.Task{
			ID:        uuid.NewString(), // generate uuid
			Name:      params.Args["name"].(string),
			Status:    model.StatusCreated, // default status: created
			CreatedAt: now,
			UpdatedAt: now,
		}
		if status, ok := params.Args["status"].(int); ok {
			task.Status = status
		}
		task.StatusStr = task.NumToStatus()
		err := task.Save()
		if err != nil {
			return nil, err
		}
		return task, nil
	},
}

var updateTask = graphql.Field{
	Type:        taskType,
	Description: "Update task by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: taskStatusEnum,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		now := time.Now()
		id := params.Args["id"].(string) // id is defined as non-null
		name, nameOk := params.Args["name"].(string)
		status, statusOK := params.Args["status"].(int)
		task, err := model.GetTaskByID(id)
		if err != nil {
			return nil, err
		}

		if nameOk {
			task.Name = name
		}
		if statusOK {
			task.Status = status
			task.StatusStr = task.NumToStatus()
		}
		task.UpdatedAt = now
		if err = task.Save(); err != nil {
			return nil, err
		}
		return task, nil
	},
}

var deleteTask = graphql.Field{
	Type:        taskType,
	Description: "Delete task by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(string)
		task, err := model.GetTaskByID(id)
		if err != nil {
			return nil, err
		}
		if err = task.Delete(); err != nil {
			return nil, err
		}
		return task, nil
	},
}
