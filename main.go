package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	// MYSQL接続

}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	// hard - api
	//	e.POST("/post", handlePOST)
	//	e.GET("/data/each", handleEachData)
	//	e.GET("/data/total", handleTotalData)

	e.GET("/test/each", handleTestEach)
	e.GET("/test/total", handleTestTotal)

	e.Logger.Fatal(e.Start(":3000"))
}
