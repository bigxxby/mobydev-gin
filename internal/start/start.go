package start

import (
	"log"
	database "project/internal/database/dataset"
	"project/internal/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func Start() {

	//init
	main, err := routes.Init()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = database.CreateMoviesTable(main.DB)
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
	router.GET("/api/movie", main.GET_Movie)
	//POST
	router.POST("/api/reg", main.POST_Reg)
	router.POST("/api/login", main.POST_Login)

	router.Run(":8080")
}
