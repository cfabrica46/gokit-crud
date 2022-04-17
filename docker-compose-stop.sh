#!/bin/bash

cd ./database-app && docker-compose down --rmi all || exit

cd ..

cd ./token-app && docker-compose down --rmi all || exit

cd ..

cd ./app && docker-compose down --rmi all || exit

docker network rm mired || exit
