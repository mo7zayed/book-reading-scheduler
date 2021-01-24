package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string

func init() {
	flag.StringVar(&port, "port", "5000", "The Port The Application Will Use To Be Served.")

	flag.Parse()
}

func main() {
	e := echo.New()

	// Root level middlewares
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Just a clean way to group routes
	{
		e.GET("/", HandleHome)
	}

	fmt.Println(
		fmt.Sprintf("Server started as http://127.0.0.1:%s", port),
	)

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", port),
		),
	)
}
