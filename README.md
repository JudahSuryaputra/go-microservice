go-microservice

## Prerequisites
- Docker and Docker Compose
- Go 1.24+ (for running locally without Docker)
- GNU Make

## Quick Start (`setup.sh`)

Start the full stack (app, db, cache, and observability stack): by executing `setup.sh` under `scripts`

## Manual

- You may check the application health by make a request to `localhost:8080/health`
- access the application log on kibana to `localhost:5601`