package router

import (
	"user-client/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()
	route.LoadHTMLGlob("./static/view/*")
	route.Static("/static/img", "./static/img")

	route.GET("/signin", func(c *gin.Context) {
		c.HTML(200, "signin.html", nil)
	})
	route.POST("/signin", handler.Signin)

	route.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})
	route.POST("/signup", handler.SignUp)

	route.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	route.POST("/query", handler.SelectUserFileAll)

	route.GET("/info", handler.UserInfo)

	route.GET("/logout", handler.Logout)
	route.POST("/rename", handler.RenameFile)

	return route
}
