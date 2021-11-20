#!/bin/bash
cd manifests/postgres
docker stop test-postgres
docker build -t test-postgres .
docker run --rm --name test-postgres -d -p 5432:5432 pasiol/postgres@sha256:bfc952e11e0202b0a55860d07b202f969da426b1447b0da21a5efdc65ea90a9f
cd ../..
docker ps | grep test-postgres