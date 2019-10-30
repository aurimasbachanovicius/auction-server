#!/usr/bin/env bash

docker-compose -f ./proxy/docker-compose.yml up -d
find ./microservices -name "docker-compose.yml" -exec docker-compose -f {} up -d \;
