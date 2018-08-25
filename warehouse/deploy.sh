#!/bin/bash
./clean.sh
docker-compose -f ./configs/fabric/docker-compose.yaml up -d