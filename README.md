# Services Work Team Search

## Description
This microservice allows searching work teams in the system using a REST API. It is built with Golang, utilizing PostgreSQL as the database.

## Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/) (1.18 or later)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Clone the Repository
To clone this project, run:
```sh
git clone https://github.com/Anggeloo/services-work-team-search.git
cd services-work-team-search
```

## Environment Variables
Create a `.env` file in the root directory and configure the required variables:

## Installation
1. Install dependencies:
   ```sh
   go mod tidy
   ```

2. Run the service:
   ```sh
   go run main.go
   ```

## Running with Docker
1. Build the Docker image:
   ```sh
   docker build -t services-work-team-search .
   ```

2. Run the container:
   ```sh
   docker run -p 8180:8180 --env-file .env services-work-team-search
   ```

3. Alternatively, use Docker Compose:
   ```sh
   docker-compose up --build
   ```

## API Documentation
This service provides API documentation via Swagger. Once the service is running, access it at:
```
http://localhost:5050/swagger/index.html
```

## Database Setup
Ensure you have a PostgreSQL database running. If using Docker, you can spin up a PostgreSQL container with:
```sh
docker run --name postgres-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=AWSPServices0099 -e POSTGRES_DB=db_work_team -p 5432:5432 -d postgres
```

## Testing
Currently, no test scripts are defined. You may add tests and run them using:
```sh
go test ./...
```

## Authors
Cadena Anggelo and Caiza Katherine
