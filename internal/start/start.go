package start

import (
	"log"
	"project/internal/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	main, err := routes.Init()
	if err != nil {
		log.Println(err.Error())
		return
	}
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
	router.GET("/api/profile", main.GET_Profile)
	router.POST("/reg", main.POST_Reg)
	router.POST("/api/login", main.POST_Login)
	log.Println("http://localhost:8080/")
	router.Run(":8080")
}
