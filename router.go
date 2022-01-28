package main

import (
	"log"
	"net/http"

	"github.com/cfabrica46/go-crud/handler"
	"github.com/cfabrica46/go-crud/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func runServer(portHTTP, portHTTPS string) {
	// gin.SetMode(gin.ReleaseMode)
	httpsRouter := setupRouterHTTPS()
	httpRouter := setupRouterHTTP(portHTTPS)

	go httpRouter.Run(":" + portHTTP)

	err := httpsRouter.RunTLS(":"+portHTTPS, "server.crt", "server.key")
	if err != nil {
		log.Println(err)
	}
}

func setupRouterHTTPS() (r *gin.Engine) {
	r = gin.Default()

	setCors(r)

	s := r.Group("/api/v1")
	s.GET("/users", handler.GetAllUsers)
	{
		getuserFromBody := s.Group("/")
		getuserFromBody.Use(middleware.GetUserFromBody)
		{
			getuserFromBody.POST("/signin", handler.SignIn)
			getuserFromBody.POST("/signup", handler.SignUp)
		}

		getuserFromToken := s.Group("/")
		getuserFromToken.Use(middleware.GetUserFromToken)
		{
			getuserFromToken.GET("/user", handler.Profile)
			getuserFromToken.DELETE("/user", handler.DeleteUser)
			getuserFromToken.HEAD("/logout", handler.LogOut)
		}
	}
	return
}

func setupRouterHTTP(portHTTPS string) (r *gin.Engine) {
	r = gin.Default()
	setCors(r)

	r.Any("/*path", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "https://localhost:"+portHTTPS+c.Request.RequestURI)
	})
	return
}

func setCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
}
