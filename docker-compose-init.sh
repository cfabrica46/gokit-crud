#!/bin/bash

# docker-compose down --rmi all

cd ./database-app && docker-compose up --build -d

cd ..

cd ./token-app && docker-compose up --build -d

cd ..

cd ./app && docker-compose up --build -d
