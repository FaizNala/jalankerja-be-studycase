package container

import (
	"jk-api/api/controllers/v1/handlers"
)

type AppContainer struct {
	ProjectHandler *handlers.ProjectHandler
	SquadHandler *handlers.SquadHandler
	TaskHandler *handlers.TaskHandler
	UserHandler *handlers.UserHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		ProjectHandler: InitProjectContainer(),
		SquadHandler: InitSquadContainer(),
		TaskHandler: InitTaskContainer(),
		UserHandler: InitUserContainer(),
	}
}
