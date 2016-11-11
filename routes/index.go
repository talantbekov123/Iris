package routes

func IndexHandler(c *iris.Context)  {
    c.Write("Hello")
}
