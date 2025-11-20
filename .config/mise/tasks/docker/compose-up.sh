#!/usr/bin/env bash
#MISE sources=["docker-compose.yml"]
#MISE description="`docker compose up -d` but with the correct file flag"

source "${MISE_PROJECT_ROOT}/.config/mise/tasks/env.sh"
"${MISE_PROJECT_ROOT}/.config/mise/tasks/docker/compose.sh" up -d