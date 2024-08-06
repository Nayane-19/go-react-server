# My Go Application

## Description

This repository contains a Go application that uses Docker Compose for container management, Tern for dependency analysis, and SQLC for generating SQL code from queries. The goal is to provide an efficient development environment and a solid foundation for Go application development.

## Prerequisites

Before you start, make sure you have the following software installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Tern](https://github.com/tern-tools/tern)
- [SQLC](https://sqlc.dev/docs/intro/)

## Project Structure

- **cmd/**: Contains the application code.
- **internal/**: Application code that should not be exposed publicly.
- **pkg/**: Reusable code across multiple packages.
- **scripts/**: Auxiliary scripts for development and automation.
- **sql/**: SQL scripts and SQLC configuration files.
- **docker-compose.yml**: Docker Compose configuration file.
- **Dockerfile**: File for building the Docker image of the application.
- **go.mod**: Go module dependency management file.
- **go.sum**: Go checksum file for dependencies.
- **README.md**: This documentation file.

## Setting Up the Environment

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. **Set Up Docker Compose**

   Docker Compose will configure the necessary services, such as the database.

   ```bash
   docker-compose up -d
   ```

   This will create and start the containers defined in the `docker-compose.yml`.

3. **Set Up SQLC**

   SQLC generates Go code from SQL files in the `sql/` directory. Ensure that the SQLC configuration file (`sqlc.yaml`) is correctly set up.

   To generate the code, run:

   ```bash
   sqlc generate
   ```

4. **Install Dependencies**

   Install the Go dependencies:

   ```bash
   go mod tidy
   ```

5. **Run the Application**

   Build and run the application:

   ```bash
   go run cmd/main.go
   ```

## Useful Scripts

- **Up Docker Compose**: `docker-compose up -d`
- **Down Docker Compose**: `docker-compose down`
- **Generate SQLC Code**: `sqlc generate`
- **Run Tests**: `go test ./...`

## Contributing

1. Fork the repository.
2. Create a branch for your feature: `git checkout -b my-feature`
3. Make your changes and commit: `git commit -am 'Add new feature'`
4. Push to the remote repository: `git push origin my-feature`
5. Create a Pull Request.