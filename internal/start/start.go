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
	//test data
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
	//ui
	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/*")
	router.Static("/static", "./ui/static")

	// HTML
	htmlRoutes := router.Group("/")
	{
		htmlRoutes.GET("/", main.GET_HTML_Index)
		htmlRoutes.GET("/reg", main.GET_HTML_Reg)
		htmlRoutes.GET("/login", main.GET_HTML_Login)
		htmlRoutes.GET("/create/movie", main.GET_HTML_Movie)
	}

	// API
	apiRoutes := router.Group("/api")
	{
		// movies
		movies := apiRoutes.Group("/movies")
		{
			movies.GET("/", main.GET_Movies)
			movies.GET("/:id", main.GET_Movie)
			movies.GET("/seasons/:id", main.GET_Season)
			movies.GET("/seasons/episodes/:id", main.GET_Episode)

			moviesAdmin := movies.Group("/")
			{
				moviesAdmin.Use(main.AuthMiddleware())

				moviesAdmin.POST("/", main.POST_Movie)        // admin
				moviesAdmin.DELETE("/:id", main.DELETE_Movie) // admin
				moviesAdmin.PUT("/:id", main.PUT_Movie)       // admin
			}
		}

		// profile
		profile := apiRoutes.Group("/profile")
		{
			profile.Use(main.AuthMiddleware())
			profile.GET("/", main.GET_Profile)
			// TODO ETC
		}

		// trends
		trends := apiRoutes.Group("/trends")
		{
			trends.GET("/:id", main.GET_Trend)
			trends.GET("/", main.GET_Trends)
		}

		// reg/log
		auth := apiRoutes.Group("/")
		{
			auth.POST("/reg", main.POST_Reg)
			auth.POST("/login", main.POST_Login)
		}

		// favorites
		favorites := apiRoutes.Group("/favorites")
		{
			favorites.GET("/", main.GET_Favorites)
			favorites.POST("/:id", main.POST_Favorite)
			favorites.DELETE("/:id", main.DELETE_Favorite)
			favorites.DELETE("/clear/", main.DELETE_Favorites)
		}
	}

	router.Run(":8080")
}
