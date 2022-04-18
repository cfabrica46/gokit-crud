#!/bin/bash

# kill Containers

cd ./database-app || exit
docker-compose kill
cd ..

cd ./token-app || exit
docker-compose kill
cd ..

cd ./app || exit
docker-compose kill
cd ..

# remove images/network/ocultContainers

cd ./database-app || exit
docker-compose down --rmi all
cd ..

cd ./token-app || exit
docker-compose down --rmi all
cd ..

cd ./app || exit
docker-compose down --rmi all
