package start

import (
	"log"
	"project/internal/routes"

	"github.com/gin-gonic/gin"
)

func Start() {

	//init
	main, err := routes.Init()
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
	//POST
	router.POST("api/reg", main.POST_Reg)
	router.POST("/api/login", main.POST_Login)
	router.Run(":8080")
}
