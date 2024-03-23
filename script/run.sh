#!/bin/bash

# Set the port to listen on.
PORT=${HTTP_LISTEN_ADDRESS:-3000}
PORT=${PORT/:/}

# Stop and remove the container if it exists.
docker rm -f web >/dev/null 2>&1

# Run the container with the environment variables from the .env file.
docker run -d --rm --name web --env-file .devcontainer/.env -p ${PORT}:${PORT} web:latest