# Financial Assistance Scheme Management System

### Overview + Tech Stack
This project is a simple REST API server written in Go for managing financial assistance schemes for needy individuals and families.

The server is built on go-chi, and uses GORM to interface with a PostgreSQL database.

1. [Local Development](#local-development)
2. [Project Structure](#project-structure)
3. [Database Schema](#database-schema)
4. [Deployment](#deployment)

### Local Development

**Prequisites**
1. Go 1.22. Download and install Go by following the instructions [here](https://go.dev/doc/install).
2. Docker compose. The easiest way to get docker compose is via installing [Docker Desktop](https://docs.docker.com/get-started/get-docker/), which includes Docker Engine and Docker CLI that are required by Compose.

#### Getting Started
1. Make a copy of the `.env.example` file as `.env`, and add in the `DB_USERNAME` and `DB_PASSWORD` values.

2. Spin up the postgres container
```
docker compose up -d
```
This will also create up a pgadmin container, accessible via http://localhost:5050/

3. Database operations (migrate, seed) are provided by the migration tool at `/cmd/database/main.go`.
To migrate the database, run 
```
make migrateDB
```
For now, this will apply all migrations by default.

4. Optionally, seed the database
```
make seedDB
```
5. Start the server
```
make run
```

### Project Structure
TODOs

### Database Schema
TODO
