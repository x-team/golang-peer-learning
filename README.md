
## Running the Application

### Docker

To run the application you can run a simple `docker-compose up --build` and it will spin up the PostgreSQL container and the Go API. 

### Taskfile

You can run the project use the tasks specified at the Taskfile.yml.

- Installation: https://taskfile.dev/installation/
- In your terminal run the `task run` command

### Rebuilding the Application

After performing changes in the code you can run the `task rebuild` command and it will rebuild the Go API docker container with the latest changes.

### Routes

The routes are defined under the `transport/http` folder. To check the available routes for this project go to [transport/http/handler.go](https://github.com/x-team/golang-peer-learning/blob/main/cmd/server/transport/http/handler.go) and look at the `mapRoutes` method

#### Authentication

Some routes require authentication where you need to send an Authorization header with a bearer token.

Steps to authenticate:
- Go to the docker-compose.yml file and change or copy the value that is set into the `TOKEN_SECRET` env variable
- Go to http://jwtbuilder.jamiekurtz.com/ scroll down to the bottom put the key you set into the `Key` input and click on `Create Signed JWT`
- Copy the JWT token and use it to authenticate with the Authorization header (example `Authorization: Bearer TOKEN`
