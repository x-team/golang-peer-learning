
## Running the Application

### Docker

To run the application you can run a simple `docker-compose up --build` and it will spin up the PostgreSQL container and the Go API. 

### Taskfile

You can run the project use the tasks specified at the Taskfile.yml.

- Installation: https://taskfile.dev/installation/
- In your terminal run the `task run` command

### Rebuilding the Application

After performing changes in the code you can run the `task rebuild` command and it will rebuild the Go API docker container with the latest changes.

