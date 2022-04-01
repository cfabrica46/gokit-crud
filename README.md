# gokit-crud

## Objetive
CRUD practice in golang using the gokit hexagonal architecture toolkit.

## Run App
~~~
docker-compose up
~~~

## Stop App
~~~
docker-compose down --rmi all
~~~

## Run Test
### All Tests
```make test```

### Test ./app
```make test-app```

### Test ./database-app
```make test-database```

### Test ./token-app
```make test-token```
