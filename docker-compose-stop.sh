#!/bin/bash

cd ./database-app && docker-compose down --rmi all

cd ..

cd ./token-app && docker-compose down --rmi all

cd ..

cd ./app && docker-compose down --rmi all

docker network rm mired
