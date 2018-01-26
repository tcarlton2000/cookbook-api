#/bin/bash

# Create docker-compose.yml from template
cp docker-compose.yml.default docker-compose.yml

# Build and deploy the database container
docker-compose build database
docker-compose up -d database

# Retrieve the database IP Address
export APP_DB_HOST=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' cookbookapi_database_1)

# Build and deploy the API container
docker-compose build api
docker-compose up -d api

# Retrieve the API IP Address
export COOKBOOK_API_HOST=http://$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' cookbookapi_api_1):8080

# Run regression tests
go test -v ./tests/...