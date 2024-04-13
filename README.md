# Tasks List go
## _A server side implementation for your tasks_

The tools and features that are in this project are as follows

- Go
- Swagger
- gin-gonic/gin
- JWT authentication
- Rate limitation
- CSRF protection
- CORS 
- MySQL
- Prometheus
- Grafana

## How does it work?
A basic implementation of Tasks list. You can create your own user (or choose one that is already created when you setup the project) and create tasks that are only yours. You can view all the tasks that are in the project assuming that all the users registered are in the same organization.

## Setup instructions
- You need to have Go installed on your system. You can check the website for Go to download according to your OS.
- Then you need to install Prometheus and Grafana. You can skip this part if you are not interested in metrics.
- MySQL is used as the database for this project. Download and install it for your system accordingly. 
- golang-migrate is used for migration so install that as well
- also, create 2 databases in your system with the names: `tasks_list_go` and `test_tasks_list_go`

## How to download Promethus and Grafana?
- You can follow the instructions on their website and download it.
- The only changes you need to make is that for Prometheus, you need to change the targets key to `localhost:8080` from `localhost:9090` so that it targets our application
- [Prometheus](https://prometheus.io/download/#prometheus) 
- [Grafana](https://grafana.com/docs/grafana/latest/getting-started/get-started-grafana-prometheus/)
- Then you need to configure golang-migrate. Instructions to download are on their website.
- [golang-migrate](https://github.com/golang-migrate/migrate)
 
## How to start the project?
- Setup the DB first
    - Navigate to /db/migrations folder
    - use the following line to create your dev and test DB 
    ```
    migrate -database "mysql://[dbUser]:[dbPassword]@tcp([dbHost]:[dbPort])/[dbName]" -path . up
    ```
    - Doing so would make 2 DB's for you. tasks_list_go (dev DB) and test_tasks_list_go (test DB)
    - You would have a dummy email already their for you along with some tasks created already
    - email and password: admin@admin.com

- Run the code with command: `go run main.go` in root folder
- Swagger instructions
    - You would need your JWT token and CSRF token to login the app and use its features
    - to get your CSRF token, route to: `GET /protected` and copy the code
    - now, you can either Register or Login  yourself
    - logging in would give you the JWT that is valid for 1 hour
    - copy your JWT token and click on either the lock icon on any of the routes or the Authorize button on the top right corner
    - be sure to add the JWT in this format `Bearer [your JWT token]` for Swagger
    - P.S. unlike Swagger, if you are using Postman, you need to go to Authorization tab and add your Bearer token only and not the `Bearer` string

- Prometheus instructions
    - Assuming you did everything that is stated above, navigate to: `localhost:9090` to view the Prometheus page
    - P.S. for windows, we run the .exe file for Prometheus wherever you downloaded the zip file and extracted it
    - in the main screen, you can choose any of following key to view your metrics
    - details for the keys are as follows
        - `get_all_tasks_list_request_count`: No. of time GET /api/v1/tasks was called
        - `get_all_user_tasks_list_request_count`: No. of time GET /api/v1/userTasks was called
        - `get_task_request_count`: No. of time GET /api/v1/task/:id was called
        - `create_task_request_count`: No. of time POST /api/v1/tasks was called
        - `edit_task_request_count`: No. of time PATCH /api/v1/editTask was called
        - `edit_task_status_request_count`: No. of time PATCH /api/v1/editTaskStatus was called
        - `delete_task_request_count`: No. of time DELETE /api/v1/deleteTask/:id was called
        - `register_user_request_count`: No. of time POST /api/v1/register was called
        - `get_user_request_count`: No. of time POST /api/v1/user was called
        - `login_request_count`: No. of time POST /api/v1/login was called
        - `edit_user_request_count`: No. of time PATCH /api/v1/editUser was called
    
- Grafana instructions
    - Assuming you did everything that is stated above, navigate to: `localhost:3000` to view the Grafana page
    - P.S. for windows, we run the .exe file for Grafana wherever you downloaded the zip file and extracted it
    - on the main screen you can see `Add your first data source`, click it and select `Prometheus`
    - give it a name
    - set the `Prometheus server URL` to `http://localhost:9090`
    - navigate to the bottom and select `Save and Test`
    - once complete, navigate to the homepage
    - on the left you can see `Dashboards`, click it and then click `New Dashboard`
    - select your datasource for that you created
    - here you can see visualisation of your data
    - add any of the metric that are listed above, set the label to `instance` and value to `localhost:9090
    - P.S. you MIGHT need to refresh Prometheus and Grafana once or twice for the changes to reflect
    
## Error codes
- `000x1`: Error querying all tasks
- `000x2`: Error scanning rows for all tasks
- `000x3`: Error querying all user tasks
- `000x4`: Error scanning rows for all user tasks
- `000x5`: No record found when fetching a single task
- `000x6`: Error when creating a task
- `000x7`: Error editing a task
- `000x8`: Error editing a task status
- `000x9`: Error deleting a task
- `000x10, 000x11 and 000x12`: Error parsing request body for Create task, Edit task and Edit Task status
- `000x20`: No record found when fetching User
- `000x21`: Error hashing user password when Registering and Editing User
- `000x22`: Error inserting User into DB when Registering
- `000x23`: Error editing user
- `000x24`: Invalid request params for fetching user
- `000x25`: Invalid request params during login
- `000x26`: Invalid credentials during login
- `000x27`: Invalid request params when editing user
- `000x28 and 000x29`: Invalid email format when creating and editing a user
- `000x30, 000x31 and 000x32`: Invalid email format when fetching User, login and editing user
- `000x33`: Email already taken when editing user
- `000x74`: Authorinzation header not present
- `000x75`: Invalid JWT
- `000x76`: Expired JWT