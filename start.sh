#!/bin/sh

docker run --name runestones-db -e POSTGRES_PASSWORD=runestones1234 -p 5432:5432 -d postgres