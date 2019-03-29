#!/usr/bin/env bash
set -ex
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql


go get -u -v ./...