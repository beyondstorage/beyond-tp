package graphql

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single task by id
			   http://localhost:7436/graphql?query={task(id:1){name,status,created_at,updated_at}}
			*/
			"task": &getTask,
			/* Get (read) task list
			   http://localhost:8080/graphql?query={tasks{name,status,created_at,updated_at}}
			*/
			"tasks": &getTaskList,
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/* Create new task item
			http://localhost:8080/graphql?query=mutation+_{create_task(name:"first create",status:"running"){id,name,status,created_at}}
			*/
			"create_task": &createTask,

			/* Update task by id
			   http://localhost:8080/graphql?query=mutation+_{update_task(id:"1",name:"test_update"){id,name,status,created_at}}
			*/
			"update_task": &updateTask,

			/* Delete task by id
			   http://localhost:8080/graphql?query=mutation+_{delete_task(id:"1"){id,name,status,created_at}}
			*/
			"delete_task": &deleteTask,
		},
	})
