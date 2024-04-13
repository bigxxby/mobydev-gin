package start

import (
	"log"
	database "project/internal/database/dataset"
	"project/pkg/middleware"
	"project/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//init
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		return
	}
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
	apiRoutes.Use(middleware.AuthMiddleware())
	{
		// movies
		movies := apiRoutes.Group("/movies")
		{
			movies.GET("/", main.MoviesRoute.GET_Movies)
			movies.GET("/:id", main.MoviesRoute.GET_Movie)
			movies.GET("/seasons/:id", main.SeasonsRoute.GET_Season)
			movies.GET("/seasons/episodes/:id", main.EpisodesRoute.GET_Episode)

			moviesAdmin := movies.Group("/")
			{
				moviesAdmin.POST("/", main.MoviesRoute.POST_Movie)        // admin
				moviesAdmin.DELETE("/:id", main.MoviesRoute.DELETE_Movie) // admin
				moviesAdmin.PUT("/:id", main.MoviesRoute.PUT_Movie)       // admin
			}
		}

		// profile
		profile := apiRoutes.Group("/profile")
		{
			profile.GET("/", main.UsersRoute.GET_Profile)
			// TODO ETC
		}

		// trends
		trends := apiRoutes.Group("/trends")
		{
			trends.GET("/:id", main.TrendsRoute.GET_Trend)
			trends.GET("/", main.TrendsRoute.GET_Trends)
		}

		// reg/log
		auth := apiRoutes.Group("/")
		{
			auth.POST("/signUp", main.AuthRoute.POST_SignUp)
			auth.POST("/signIn", main.AuthRoute.POST_SignIn)
		}

		// favorites
		favorites := apiRoutes.Group("/favorites")
		{
			favorites.GET("/", main.FavoritesRoute.GET_Favorites)
			favorites.POST("/:id", main.FavoritesRoute.POST_Favorite)
			favorites.DELETE("/:id", main.FavoritesRoute.DELETE_Favorite)
			favorites.DELETE("/clear/", main.FavoritesRoute.DELETE_Favorites)
		}
	}

	router.Run(":8080")
}
