#!/usr/bin/env zsh

go build -ldflags "-X 'main.GitTag=$(git describe --tags)' -X 'main.Timestamp=$(date -u)'"

