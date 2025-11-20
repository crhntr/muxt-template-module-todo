#!/usr/bin/env bash
#MISE description="start the docker compose"

COMPOSE_FILEPATH="${MISE_PROJECT_ROOT}/.config/docker-compose.yml"

docker compose --file="${COMPOSE_FILEPATH}" "$@"
