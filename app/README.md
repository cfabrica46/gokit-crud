# database-app

## Run App
~~~
docker-compose up
~~~

## Stop App
~~~
docker-compose down --rmi all
~~~

## Run Test
~~~
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
~~~
