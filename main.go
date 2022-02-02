package main

import (
	"gint"
	"net/http"
)

func main() {
	r := gint.New()
	r.GET("/", func(c *gint.Context) {
		c.HTML(http.StatusOK, "<h1>Index Home</h1>")
	})

	v1 := r.Group("/v1")
	v1.GET("/", func(c *gint.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	v1.GET("/hello", func(c *gint.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2 := r.Group("/v2")
	v2.GET("/hello/:name", func(c *gint.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v2.POST("/login", func(c *gint.Context) {
		c.JSON(http.StatusOK, gint.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
