#!/usr/bin/env bash
#MISE description="run the server"
#MISE depends=["docker:compose-up"]

go -C "${MISE_PROJECT_ROOT}/cmd/server" run .