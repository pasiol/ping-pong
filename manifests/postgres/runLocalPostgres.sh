#!/bin/bash
docker build -t test-postgres .
docker run --rm --name test-postgres -d -p 5432:5432 test-postgres:latest