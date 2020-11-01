# Todo Server

## Installation

First, pull the project into your working environment:

```zsh
git pull git@github.com/mayoz/todo-server
```

## Development Environment

The following requirements are assumed to be installed on your local environment:

- [Go](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/downloads/)

You can use following command for running app:

```zsh
go run {service-path}
```

Example usage for "api" service:

```zsh
go run ./cmd/api
```

Or, you can use Docker for your development environment.
Please, first check the following requirements on your computer:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

Check the `.env` file and run:

```zsh
docker-compose up --build -d
```

You should be able to access the app on [`localhost:8080`](http://localhost:8080).

#### Testing

```shell script
make test
```

If you need coverage you can run the following command:

```shell script
make test-cover
```
