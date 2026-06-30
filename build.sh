#!/bin/bash
set -e

echo "Installing dependencies..."
go mod tidy \
pnpm --dir frontend install

echo "Building the frontend..."
pnpm --dir frontend run build

echo "Compiling..."
CGO_ENABLED=0 go build -o ./bin/filecrow -trimpath -ldflags='-s -w'

echo "Done!"

