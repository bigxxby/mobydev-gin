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

	err = database.DropTables(main.DB.Database)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = database.CreateTables(main.DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = database.CreateTestData(main.DB)
	if err != nil {
		log.Println(err.Error())
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/*")
	router.Static("/static", "./ui/static")

	//HTML
	router.GET("/", main.GET_HTML_Index)
	router.GET("/reg", main.GET_HTML_Reg)
	router.GET("/login", main.GET_HTML_Login)
	router.GET("/create/movie", main.GET_HTML_Movie)

	//GET
	router.GET("/api/profile", main.GET_Profile)

	router.GET("/api/movies", main.GET_Movies)
	router.GET("/api/movies/:id", main.GET_Movie)

	router.GET("/api/movies/season/:id", main.GET_Season)
	router.GET("/api/movies/episode/:id", main.GET_Episode)

	router.GET("/api/trends/:id", main.GET_Trend)
	router.GET("/api/trends", main.GET_Trends)

	//POST
	router.POST("/api/reg", main.POST_Reg)
	router.POST("/api/login", main.POST_Login)
	router.POST("/api/movies", main.POST_Movie)

	//DELETE
	router.DELETE("/api/movies/:id", main.DELETE_Movie)

	//PUT
	router.PUT("/api/movies/:id", main.PUT_Movie)

	router.Run(":8080")
}
