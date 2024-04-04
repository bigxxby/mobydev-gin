package main

import (
	"log"
	"project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// creating main manager
	main, err := handlers.Init()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// err = main.DB.CreateProjectsTable() //////////////////////////////
	if err != nil {
		log.Println(err.Error())
	}

	// main.DB.CreateUsersTable() ///////////////////////////////
	// main.DB.CreateAdmin()

	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/*")
	router.Static("/static", "./ui/static")
	//routes
	router.GET("/", main.GET_Index)
	router.GET("/reg", main.GET_Reg)
	router.GET("/login", main.GET_Login)
	router.GET("/api/profile/:sessionId", main.GET_Profile)
	router.POST("/reg", main.POST_Reg)
	router.POST("/api/login", main.POST_Login)
	router.POST("/api/logout", main.POST_Logout)
	log.Println("http://localhost:8080/")
	router.Run(":8080")
}
