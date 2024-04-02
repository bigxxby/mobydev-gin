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
	err = main.DB.CreateProjectsTable() //////////////////////////////
	if err != nil {
		log.Println(err.Error())
	}
	main.DB.CreateUsersTable() ///////////////////////////////
	// main.DB.CreateAdmin()
	router := gin.Default()
	// static templates
	router.LoadHTMLGlob("ui/templates/*")
	// static css
	router.Static("/static", "./ui/static")
	// -> /
	router.GET("/", main.Index)
	// -> /reg
	router.GET("/reg", main.RegHandler)
	router.POST("/reg", main.RegHandler)
	// -> /login
	router.GET("/login", main.LogHandler)
	router.POST("/login", main.LogHandler)
	router.POST("/get-profile", main.GetProfile)
	router.POST("/logout", main.Logout)
	// router.GET("/", handlers.ProfileHandler)
	// router.GET("/", handlers.LogoutHandler)
	// router.GET("/", handlers.ProjectsHandler)
	// router.GET("/", handlers.CreateProjectsHandler)
	log.Println("http://localhost:8080/")
	router.Run(":8080")
}

// mHandler := handlers.Init_handler()
// if mHandler == nil {
// 	return
// }
// err := mHandler.Data.CreateProjectsTable()
// err = mHandler.Data.CreateUsersTable()
// mHandler.Data.CreateAdmin()
// if err != nil {
// 	log.Println("Error creating tables : ", err.Error())
// 	return
// }
// fileServer := http.FileServer(http.Dir("./ui/static/"))
// http.Handle("/static/", http.StripPrefix("/static/", fileServer))
// http.HandleFunc("/", mHandler.IndexHandler)
// http.HandleFunc("/reg", mHandler.RegHandler)
// http.HandleFunc("/login", mHandler.LogHandler)
// http.HandleFunc("/profile", mHandler.ProfileHandler)
// http.HandleFunc("/logout", mHandler.LogoutHandler)
// http.HandleFunc("/projects", mHandler.ProjectsHandler)
// http.HandleFunc("/createProject", mHandler.CreateProjectsHandler)
// http.HandleFunc("/getProjects", mHandler.GetProjectsHandler)
// http.HandleFunc("/updateProfile", mHandler.UpdateProfileHandler)
// http.HandleFunc("/admin", mHandler.AdminHandler)
// http.HandleFunc("/getAllUsers", mHandler.GetAllUsersHandler)
// http.HandleFunc("/updateUserAdmin", mHandler.UpdateUserAdminHandler)
// http.HandleFunc("/deleteUser/", mHandler.DeleteUserAdminHandler)
