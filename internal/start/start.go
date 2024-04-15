package start

import (
	"log"
	"project/internal/database/datasets"
	"project/pkg/middleware"
	"project/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		return
	}
	//init main
	main, err := routes.Init()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// test data
	err = datasets.InitDatasets(main.DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
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
			movies.POST("/", main.MoviesRoute.POST_Movie)        // admin
			movies.DELETE("/:id", main.MoviesRoute.DELETE_Movie) // admin
			movies.PUT("/:id", main.MoviesRoute.PUT_Movie)       // admin
		}
		//seasons
		seasons := apiRoutes.Group("/seasons")
		{
			seasons.GET("/seasons/:id", main.SeasonsRoute.GET_Season)
		}
		//episodes
		episodes := apiRoutes.Group("/seasons")
		{
			episodes.GET("/:id", main.EpisodesRoute.GET_Episode)
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
			favorites.GET("/", main.FavoritesRoute.GET_Favorites)             //get fav of CURRENT USER
			favorites.POST("/:id", main.FavoritesRoute.POST_Favorite)         //get fav by id  of CURRENT USER
			favorites.DELETE("/:id", main.FavoritesRoute.DELETE_Favorite)     //delete fav of CURRENT USER
			favorites.DELETE("/clear/", main.FavoritesRoute.DELETE_Favorites) //delete all fav of CURRENT USER
		}
		// categories
		categories := apiRoutes.Group("/categories")
		{
			categories.GET("/", main.CategoriesRoute.GET_Categories)
			categories.GET("/:id", main.CategoriesRoute.GET_Category)
			categories.POST("/", main.CategoriesRoute.POST_Category)        ////admin
			categories.PUT("/:id", main.CategoriesRoute.PUT_Category)       ////admin
			categories.DELETE("/:id", main.CategoriesRoute.DELETE_Category) ////admin
		}
		// genres
		genres := apiRoutes.Group("/genres")
		{
			genres.GET("/", main.GenreRoute.GET_Genres)

			genres.GET("/:id", main.GenreRoute.GET_Genre)
			genres.POST("/", main.GenreRoute.POST_Genre)        ////admin
			genres.DELETE("/:id", main.GenreRoute.DELETE_Genre) ////admin
			genres.PUT("/:id", main.GenreRoute.PUT_Genre)       ////admin

		}
		//age
		age := apiRoutes.Group("/ageCategories")
		{
			age.GET("/", main.AgeRoute.GET_AgeCategories)

			age.GET("/:id", main.AgeRoute.GET_AgeCategory)
			age.POST("/", main.AgeRoute.POST_AgeCategory) ////admin
			// age.DELETE("/:id", main.GenreRoute.DELETE_Genre) ////admin
			// age.PUT("/:id", main.GenreRoute.PUT_Genre)       ////admin
		}

	}

	router.Run(":8080")
}
