package routes

import "github.com/prometheus/client_golang/prometheus"

// Could've created an array and initialise the counters in a loop
// That would result in accessing the counters in user and task routes by array indexing
// Creating separate variables makes the code a bit more readable and easier to understand
// -----------------------------------------------------------
// task related request counters
var tasksRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_all_tasks_list_request_count",
		Help: "No. of time GET /api/v1/tasks was called",
	},
)

var userTasksRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_all_user_tasks_list_request_count",
		Help: "No. of time GET /api/v1/userTasks was called",
	},
)

var taskRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_task_request_count",
		Help: "No. of time GET /api/v1/task/:id was called",
	},
)

var createTasksRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "create_task_request_count",
		Help: "No. of time POST /api/v1/tasks was called",
	},
)

var editTaskRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "edit_task_request_count",
		Help: "No. of time PATCH /api/v1/editTask was called",
	},
)

var editTaskStatusRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "edit_task_status_request_count",
		Help: "No. of time PATCH /api/v1/editTaskStatus was called",
	},
)

var deleteTasksRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "delete_task_request_count",
		Help: "No. of time DELETE /api/v1/deleteTask/:id was called",
	},
)

// -----------------------------------------------------------
// user related request counters
var registerUserRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "register_user_request_count",
		Help: "No. of time POST /api/v1/register was called",
	},
)

var getUserRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "get_user_request_count",
		Help: "No. of time POST /api/v1/user was called",
	},
)

var loginRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "login_request_count",
		Help: "No. of time POST /api/v1/login was called",
	},
)

var editUserRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "edit_user_request_count",
		Help: "No. of time PATCH /api/v1/editUser was called",
	},
)

// -----------------------------------------------------------
// function to register the counters so that we can display it on Prometheus and Grafana
func setupPrometheusCounters() {
	prometheus.MustRegister(tasksRequestCounter)
	prometheus.MustRegister(userTasksRequestCounter)
	prometheus.MustRegister(taskRequestCounter)
	prometheus.MustRegister(createTasksRequestCounter)
	prometheus.MustRegister(editTaskRequestCounter)
	prometheus.MustRegister(editTaskStatusRequestCounter)
	prometheus.MustRegister(deleteTasksRequestCounter)
	prometheus.MustRegister(registerUserRequestCounter)
	prometheus.MustRegister(getUserRequestCounter)
	prometheus.MustRegister(loginRequestCounter)
	prometheus.MustRegister(editUserRequestCounter)
}
