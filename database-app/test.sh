# GetAllUsers
curl -XGET localhost:7070/users

# GetUserByID
# curl -XGET localhost:7070/user/1

# GetUserByUsernameAndPassword
# curl -XGET -d'{"username":"cesar","password":"01234"}' localhost:7070/user/username_password

# GetIDByUsername
# curl -XGET localhost:7070/id/cesar

# Insert User
# curl -XPOST -d'{"username":"arturo","password":"nava","email":"arthurnavah@gmail.com"}' localhost:7070/user

# DeleteUserByUsername
# curl -XDELETE -d'{"username":"arturo","password":"nava","email":"arthurnavah@gmail.com"}' localhost:7070/user

