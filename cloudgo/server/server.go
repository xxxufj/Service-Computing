package server

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Serve() {
	m := martini.Classic()

	// render html templates from templates directory
	m.Use(render.Renderer())

	// use pictures in assets directory
	m.Use(martini.Static("assets"))

	// This will set the Content-Type header to "application/html; charset=ISO-8859-1"
	m.Get("/", func(r render.Render) {
		//render 'login.tmpl'
		r.HTML(200, "login", map[string]interface{}{})
	})

	m.Post("/", binding.Bind(User{}), func(u User, r render.Render) {
		p := User{Username: u.Username, Password: u.Password}
		// render 'info.tmpl', and add information of user to the webpage
		r.HTML(200, "info", map[string]interface{}{"user": p})
	})

	// start listening
	m.Run()
}
