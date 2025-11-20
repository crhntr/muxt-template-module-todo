#!/usr/bin/env bash
#MISE description="source this to get environment configuration"

COMPOSE_FILEPATH=".config/docker-compose.yml"

POSTGRES_USER="$(yq '.services.postgres.environment.POSTGRES_USER' $COMPOSE_FILEPATH)"
POSTGRES_PASSWORD="$(yq '.services.postgres.environment.POSTGRES_PASSWORD' $COMPOSE_FILEPATH)"
POSTGRES_DB="$(yq '.services.postgres.environment.POSTGRES_DB' $COMPOSE_FILEPATH)"

export DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"
