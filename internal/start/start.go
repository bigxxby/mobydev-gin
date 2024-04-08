package start

import (
	"log"
	database "project/internal/database/dataset"
	"project/internal/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//init
	main, err := routes.Init()
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = database.DropTables(main.DB.Database) ///////////////
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = database.CreateTables(main.DB)
	if err != nil {
		return
	}
	err = database.CreateTestData(main.DB) /////////////
	if err != nil {
		log.Println(err.Error())
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/*")
	router.Static("/static", "./ui/static")

	//GET
	router.GET("/", main.GET_Index)
	router.GET("/reg", main.GET_Reg)
	router.GET("/login", main.GET_Login)
	router.GET("/api/profile", main.GET_Profile)
	router.GET("/api/project", main.GET_Project)
	router.GET("/api/project/:id", main.GET_ProjectById)
	router.GET("/forgot", main.GET_Forgot)
	//POST
	router.POST("/api/reg", main.POST_Reg)
	router.POST("/api/forgot", main.POST_Forgot)
	router.POST("/api/login", main.POST_Login)

	router.Run(":8080")

}
