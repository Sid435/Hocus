package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sid/Hocus/api"
	"github.com/sid/Hocus/hocus"
)

func main() {
	var a any
	a = false
	b := false
	fmt.Println(a == b)
	db, err := hocus.New()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(db)

	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusInternalServerError, hocus.Map{"error": err.Error()})
	}
	e.HideBanner = true
	e.POST("/api/:collname", server.HandlePostInsert)
	e.GET("/api/:collname", server.HandleGetQuery)
	log.Fatal(e.Start(":7777"))
}
