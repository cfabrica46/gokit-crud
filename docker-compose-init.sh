#!/bin/bash

docker network create mired

cd ./database-app && docker-compose up --build -d

cd ..

cd ./token-app && docker-compose up --build -d

cd ..

cd ./app && docker-compose up --build -d
