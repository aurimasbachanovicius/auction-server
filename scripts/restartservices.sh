#!/usr/bin/env bash

find ./microservices -name "docker-compose.yml" -exec docker-compose -f {} restart \;
