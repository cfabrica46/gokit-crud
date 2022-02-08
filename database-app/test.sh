# GetAllUsers
# curl -XGET localhost:8080/users

# GetUserByID
# curl -XGET localhost:8080/user/1

# GetUserByUsernameAndPassword
# curl -XGET -d'{"username":"cesar","password":"01234"}' localhost:8080/user/username_password

# GetIDByUsername
# curl -XGET localhost:8080/id/cesar

# Insert User
# curl -XPOST -d'{"username":"arturo","password":"nava","email":"arthurnavah@gmail.com"}' localhost:8080/user

# DeleteUserByUsername
# curl -XDELETE -d'{"username":"arturo","password":"nava","email":"arthurnavah@gmail.com"}' localhost:8080/user

