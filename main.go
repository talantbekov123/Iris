package main

import (
	"gopkg.in/kataras/iris.v4"
	_ "fmt"
	//"iris/routes"
)

func IndexHandler(c *iris.Context)  {
    c.Write("Hello")
}

func main() {

	iris.Get("/", IndexHandler)


	iris.Listen(":8080")
}