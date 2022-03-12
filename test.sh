#!/bin/bash

#SignUp
# response=$(curl -X POST -k https://localhost:8081/api/v1/signup -d '{"username":"cesar","password":"01234","email":"cfabrica46@gmail.com"}')

#Signin
response=$(curl -X POST -Lk http://localhost:8080/signin -d '{"username":"cesar","password":"01234"}')

token=$(echo "$response" | jq -r '.content')

echo "$token"

#ShowUsers
# curl -X GET -Lk http://localhost:8080/users

#Profile
curl -X POST -Lk http://localhost:8080/profile -d `{"token":"($token)"}`
# curl -X POST -k http://localhost:8081/api/v1/user -H "Authorization: $token"

#Delete
# curl -X DELETE -k http://localhost:8080/profile -d "{'token':$(token)}" 
# curl -X DELETE -Lk http://localhost:8080/api/v1/user -H "Authorization: $token"

