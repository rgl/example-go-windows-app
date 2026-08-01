#!/usr/bin/bash
set -euxo pipefail

export APP_VERSION="$1"

go generate ./...
