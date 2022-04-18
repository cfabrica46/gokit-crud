#!/bin/bash

# docker network create mired || exit

cd ./database-app || exit
docker-compose up --build -d
cd ..

cd ./token-app || exit
docker-compose up --build -d
cd ..

cd ./app || exit
docker-compose up --build -d
