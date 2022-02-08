# GetAllUsers
# curl localhost:8080/users

# GetUserByID
curl localhost:8080/user/1

# GetUserByUsernameAndPassword
# curl -XPOST -d'{"username":"cesar","password":"caycho"}' localhost:8080/userByUsernameAndPassword

# GetIDByUsername
# curl -XPOST -d'{"username":"cesar"}' localhost:8080/idByUsername

# Insert User
# curl -XPOST -d'{"username":"cesar","password":"caycho","email":"cfabrica46@gmail.com"}' localhost:8080/insert

# DeleteUserByUsername
# curl -XPOST -d'{"username":"cesar"}' localhost:8080/delete
