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

		//CHANGE PASSWORD LOGIC ->
		//sends verification code to the given email ->
		htmlRoutes.GET("/send-code", main.GET_HTML_SendRestoreCode) //
		//auth.POST("/send-code", main.AuthRoute.POST_Restore) // sends restore code to the users email
		//auth.POST("/verify", main.AuthRoute.POST_Verify)   // gets code from user and verifies it,  if -
		//- valid sends to the email the link with token for resetting the password and to the frontend to redirect the user,
		htmlRoutes.GET("/change-password", main.GET_ChangePassword) // if token and email is valid gives perm. to change the password
		//auth.POST("/reset-password", main.AuthRoute.POST_ResetPassword) // gets new password from user and change it

	}

	// API
	apiRoutes := router.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware())
	{
		// movies
		movies := apiRoutes.Group("/movies")
		{
			movies.GET("/", main.MoviesRoute.GET_Movies)          //?limit=<limitInt>
			movies.GET("/:id", main.MoviesRoute.GET_Movie)        // returns movie by id of movie
			movies.POST("/", main.MoviesRoute.POST_Movie)         // admin only
			movies.DELETE("/:id", main.MoviesRoute.DELETE_Movie)  // admin only
			movies.PUT("/:id", main.MoviesRoute.PUT_Movie)        // admin only
			movies.GET("/main", main.MoviesRoute.GET_Movies_MAIN) //?limit=<limitInt> (with category, age category,genre)
			movies.GET("/search", main.MoviesRoute.GET_Search)    // ?query=<searchQuery>
		}
		//seasons
		seasons := apiRoutes.Group("/seasons")
		{
			seasons.GET("/seasons/:id", main.SeasonsRoute.GET_Season) // returns all seasons (only get)
		}
		//episodes
		episodes := apiRoutes.Group("/seasons")
		{
			episodes.GET("/:id", main.EpisodesRoute.GET_Episode) // returns all episodes (only get)
		}

		// profile
		profile := apiRoutes.Group("/profile")
		{
			profile.GET("/", main.UsersRoute.GET_Profile) //get profile of current user
			profile.PUT("/", main.UsersRoute.PUT_Profile) //update profile of current user (dob , name , phone)
		}

		// trends
		trends := apiRoutes.Group("/trends") // only get
		{
			trends.GET("/:id", main.TrendsRoute.GET_Trend) //returns trend by id
			trends.GET("/", main.TrendsRoute.GET_Trends)   //returns all of the current trends
		}

		// auth
		auth := apiRoutes.Group("/")
		{
			auth.POST("/signUp", main.AuthRoute.POST_SignUp) //{email, password, role } required bindings
			auth.POST("/signIn", main.AuthRoute.POST_SignIn) //{email, password } required bindings

			//htmlRoutes.GET("/send-code", main.GET_HTML_SendRestoreCode) //-------->
			auth.POST("/send-code", main.AuthRoute.POST_SendCode) // sends restore code to the users email
			auth.POST("/verify", main.AuthRoute.POST_VerifyCode)  // gets code from user an verifies it, if valid sends to the email link for resetting the pass
			// htmlRoutes.GET("/change-password", main.GET_ChangePassword) //------->
			auth.POST("/reset-password", main.AuthRoute.POST_ResetPassword) // changes password of the user

		}

		// favorites
		favorites := apiRoutes.Group("/favorites")
		{
			favorites.GET("/", main.FavoritesRoute.GET_Favorites)             //returns fav movies of CURRENT USER
			favorites.POST("/:id", main.FavoritesRoute.POST_Favorite)         //returns fav movie by id of CURRENT USER
			favorites.DELETE("/:id", main.FavoritesRoute.DELETE_Favorite)     //delete fav movie of CURRENT USER
			favorites.DELETE("/clear/", main.FavoritesRoute.DELETE_Favorites) //delete all fav movies of CURRENT USER
		}
		// categories
		categories := apiRoutes.Group("/categories")
		{
			categories.GET("/", main.CategoriesRoute.GET_Categories)        // return all categories
			categories.GET("/:id", main.CategoriesRoute.GET_Category)       // returns category by id
			categories.POST("/", main.CategoriesRoute.POST_Category)        //admin only
			categories.PUT("/:id", main.CategoriesRoute.PUT_Category)       //admin only
			categories.DELETE("/:id", main.CategoriesRoute.DELETE_Category) //admin only
		}
		// genres
		genres := apiRoutes.Group("/genres")
		{
			genres.GET("/", main.GenreRoute.GET_Genres)         //return all genres
			genres.GET("/:id", main.GenreRoute.GET_Genre)       //return genre by id
			genres.POST("/", main.GenreRoute.POST_Genre)        ////admin
			genres.DELETE("/:id", main.GenreRoute.DELETE_Genre) ////admin
			genres.PUT("/:id", main.GenreRoute.PUT_Genre)       ////admin

		}
		//age
		age := apiRoutes.Group("/ageCategories")
		{
			age.GET("/", main.AgeRoute.GET_AgeCategories)        // return all age categories
			age.GET("/:id", main.AgeRoute.GET_AgeCategory)       //return category by id
			age.POST("/", main.AgeRoute.POST_AgeCategory)        //admin only
			age.DELETE("/:id", main.AgeRoute.DELETE_AgeCategory) //admin only
			age.PUT("/:id", main.AgeRoute.PUT_AgeCategory)       //admin only
		}

	}

	router.Run(":8080")
}
