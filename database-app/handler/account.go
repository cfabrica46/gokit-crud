package handler

/* func SignUp(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	email := c.MustGet("email").(string)

	passwordHash := newSha256([]byte(password))
	password = fmt.Sprintf("%x", passwordHash)

	err := userdb.InsertUser(username, password, email)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to insert user"})
		return
	}

	id, err := userdb.GetIDByUsername(username)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to insert user"})
		return
	}

	keyData, err := ioutil.ReadFile("server.key")
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Creating Token"})
		return
	}

	userToken, err := token.GenerateToken(id, username, email, keyData, jwt.SigningMethodHS256)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Creating Token"})
		return
	}

	err = cache.SetToken(userToken)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to set Token"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: userToken})
} */

/* func SignIn(c *gin.Context) {
	username := c.MustGet("username").(string)
	password := c.MustGet("password").(string)
	// email := c.MustGet("email").(string)

	passwordHash := newSha256([]byte(password))
	password = fmt.Sprintf("%x", passwordHash)

	user, err := userdb.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Sign In"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, structure.ResponseHTTP{Code: http.StatusUnauthorized, ErrorText: "Error User Not Found"})
		return
	}

	keyData, err := ioutil.ReadFile("server.key")
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Creating Token"})
		return
	}

	userToken, err := token.GenerateToken(user.ID, user.Username, user.Email, keyData, jwt.SigningMethodHS256)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to create token"})
		return
	}

	err = cache.SetToken(userToken)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to set Token"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: userToken})
} */

/* func LogOut(c *gin.Context) {
	userToken := c.MustGet("token").(string)

	err := cache.DeleteTokenUsingValue(userToken)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to logout"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK})
} */
