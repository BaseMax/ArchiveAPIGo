package main

import (
	"archive/apis"
	"archive/dbconection"
	"archive/service/categories"
	"archive/service/files"
	"archive/service/users"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {

	dbconection.ConnectDatabase()
	dbconection.DB.AutoMigrate(&categories.Category{})
	dbconection.DB.AutoMigrate(&users.User{})
	dbconection.DB.AutoMigrate(&files.File{})
}

func main() {

	e := echo.New()

	apis.Routes(e)

	s := fmt.Sprintf(":%s", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(s))
}
