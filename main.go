package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iam1912/student-admin/handlers"
)

func main() {
	r := gin.Default()
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("./templates/*/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		handlers.LoginHandler(c)
	})

	admin := r.Group("/admin")
	{
		admin.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", nil)
		})
		admin.GET("/index", func(c *gin.Context) {
			handlers.IndexHandler(c)
		})
		admin.POST("/search", func(c *gin.Context) {
			handlers.SearchHandler(c)
		})
		admin.GET("/add", func(c *gin.Context) {
			c.HTML(http.StatusOK, "add.html", nil)
		})
		admin.POST("/add", func(c *gin.Context) {
			handlers.AddHandler(c)
		})
		admin.POST("/delete", func(c *gin.Context) {
			handlers.DeleteHandler(c)
		})
		admin.GET("/edit", func(c *gin.Context) {
			handlers.PreviewHandler(c)
		})
		admin.POST("/edit", func(c *gin.Context) {
			handlers.EditHandler(c)
		})
	}

	r.Run(":8080")
}
