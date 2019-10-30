#!/usr/bin/env bash

find ./microservices -name "docker-compose.yml" -exec docker-compose -f {} down \;
docker-compose -f ./proxy/docker-compose.yml down
