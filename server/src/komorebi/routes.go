package komorebi

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.NotFoundHandler = http.HandlerFunc(OwnNotFound)

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetBoards",
		"GET",
		"/boards",
		BoardsGet,
	},
	Route{
		"GetStories",
		"GET",
		"/{board_name}/stories",
		GetStories,
	},
	Route{
		"BoardCreate",
		"POST",
		"/boards",
		BoardCreate,
	},
	Route{
		"BoardUpdate",
		"POST",
		"/boards/{board_id}",
		BoardUpdate,
	},
	Route{
		"AssignUsersToBoard",
		"POST",
		"/boards/{board_id}/assign_users",
		AssignUsersToBoard,
	},
	Route{
		"GetUsersFromBoard",
		"GET",
		"/boards/{board_id}/users",
		GetUsersFromBoard,
	},
	Route{
		"ClearDumpsFromBoard",
		"GET",
		"/boards/{board_id}/clear",
		ClearDumpsFromBoard,
	},
	Route{
		"BoardDelete",
		"DELETE",
		"/boards/{board_id}",
		BoardDelete,
	},
	Route{
		"GetBurndownFromBoard",
		"GET",
		"/boards/{board_id}/burndown",
		GetBurndownFromBoard,
	},
	Route{
		"ColumnGet",
		"GET",
		"/columns/{column_id}",
		ColumnGet,
	},
	Route{
		"ColumnUpdate",
		"POST",
		"/columns/{column_id}",
		ColumnUpdate,
	},
	Route{
		"ColumnMove",
		"POST",
		"/columns/{column_id}/move",
		ColumnMove,
	},
	Route{
		"TasksGetByColumn",
		"GET",
		"/columns/{column_id}/tasks",
		TasksGetByColumn,
	},
	Route{
		"ColumnCreate",
		"POST",
		"/columns",
		ColumnCreate,
	},
	Route{
		"ColumnDelete",
		"DELETE",
		"/columns/{column_id}",
		ColumnDelete,
	},
	Route{
		"StoryCreate",
		"POST",
		"/stories",
		StoryCreate,
	},
	Route{
		"StoryGet",
		"GET",
		"/stories/{story_id}",
		StoryGet,
	},
	Route{
		"StoryDodGet",
		"GET",
		"/stories/{story_id}/dods",
		StoryDodGet,
	},
	Route{
		"StoryDodUpdate",
		"POST",
		"/stories/{story_id}/dods",
		StoryDodUpdate,
	},
	Route{
		"StoryUpdate",
		"POST",
		"/stories/{story_id}",
		StoryUpdate,
	},
	Route{
		"StoryDelete",
		"DELETE",
		"/stories/{story_id}",
		StoryDelete,
	},
	Route{
		"UsersGet",
		"GET",
		"/users",
		UsersGet,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreate,
	},
	Route{
		"UserUpdate",
		"POST",
		"/users/{user_id}",
		UserUpdate,
	},
	Route{
		"UserDelete",
		"DELETE",
		"/users/{user_id}",
		UserDelete,
	},
	Route{
		"HandleWS",
		"GET",
		"/{board_name}/ws",
		HandleWs,
	},
	Route{
		"TaskCreate",
		"POST",
		"/tasks",
		TaskCreate,
	},
	Route{
		"TaskUpdate",
		"POST",
		"/tasks/{task_id}",
		TaskUpdate,
	},
	Route{
		"AssignUsersToTask",
		"POST",
		"/tasks/{task_id}/assign_users",
		AssignUsersToTask,
	},
	Route{
		"GetUsersFromTask",
		"GET",
		"/tasks/{task_id}/users",
		GetUsersFromTask,
	},
	Route{
		"TaskDelete",
		"DELETE",
		"/tasks/{task_id}",
		TaskDelete,
	},
	Route{
		"TasksGet",
		"GET",
		"/stories/{story_id}/tasks",
		TasksGet,
	},
	Route{
		"TaskGet",
		"GET",
		"/tasks/{task_id}",
		TaskGet,
	},
	Route{
		"BoardDodGet",
		"GET",
		"/{board_name}/dods",
		BoardDodGet,
	},
	Route{
		"BoardDodUpdate",
		"POST",
		"/{board_name}/dods",
		BoardDodUpdate,
	},
	Route{
		"BoardShow",
		"GET",
		"/{board_name}",
		BoardShow,
	},
	Route{
		"GetFeatureAndCreateStory",
		"GET",
		"/create_story_by_issue/{board_id}/{issue}",
		GetFeatureAndCreateStory,
	},
}
