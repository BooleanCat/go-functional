#!/usr/bin/env bash
set -euo pipefail

go vet ./...
go fmt ./...
golint --set_exit_status
ginkgo -r --race --randomizeAllSpecs --randomizeSuites .
