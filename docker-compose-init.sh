#!/bin/bash

docker network create mired || exit

cd ./database-app && docker-compose up --build -d || exit

cd ..

cd ./token-app && docker-compose up --build -d || exit

cd ..

cd ./app && docker-compose up --build -d || exit
