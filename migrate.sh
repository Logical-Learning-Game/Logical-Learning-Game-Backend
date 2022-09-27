#!/bin/sh


echo "Run DB Migration"

migrate -path migrations -database "mysql://$1" -verbose up