#!/bin/bash
go get github.com/99designs/gqlgen/cmd@v0.14.0  github.com/99designs/gqlgen/internal/imports@v0.14.0 github.com/99designs/gqlgen/internal/code@v0.14.0 github.com/99designs/gqlgen/internal/imports@v0.14.0
go get ./...
go run github.com/99designs/gqlgen generate