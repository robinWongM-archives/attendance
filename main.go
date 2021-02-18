package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"time"
)

func main() {
	app := iris.New()
	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	//
	// Need to use a custom regexp instead?
	// Easy;
	// Just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	// Start the server using a network address.
	app.Listen(":8080", iris.WithoutInterruptHandler)
}