#!/bin/bash

# kill Containers

cd ./database-app && docker-compose kill || exit

cd ..

cd ./token-app && docker-compose kill || exit

cd ..

cd ./app && docker-compose kill || exit

cd ..

# Remove Images/Network/OcultContainers

cd ./database-app && docker-compose down --rmi all || exit

cd ..

cd ./token-app && docker-compose down --rmi all || exit

cd ..

cd ./app && docker-compose down --rmi all || exit
