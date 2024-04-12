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

	// HTML
	router.GET("/", main.GET_HTML_Index)
	router.GET("/reg", main.GET_HTML_Reg)
	router.GET("/login", main.GET_HTML_Login)
	router.GET("/create/movie", main.GET_HTML_Movie)

	// movies
	router.GET("/api/movies", main.GET_Movies)              //get all movies
	router.GET("/api/movies/:id", main.GET_Movie)           // get movie by id
	router.GET("/api/movies/season/:id", main.GET_Season)   //get season by id
	router.GET("/api/movies/episode/:id", main.GET_Episode) // get episode by id

	router.POST("/api/movies", main.POST_Movie) //create movie (ADMIN ONLY)

	router.DELETE("/api/movies/:id", main.DELETE_Movie) // delete movie by id (ADMIN ONLY)

	router.PUT("/api/movies/:id", main.PUT_Movie) //change movie by id (ADMIN ONLY)

	// profile
	router.GET("/api/profile", main.GET_Profile) // get profile of CURRENT USER (auth)
	// get all profiles
	// get profile by id etc.

	// trends
	router.GET("/api/trends/:id", main.GET_Trend) // get trend by id
	router.GET("/api/trends", main.GET_Trends)    // get all current trends

	// reg/log
	router.POST("/api/reg", main.POST_Reg)     // body{email,password,role}
	router.POST("/api/login", main.POST_Login) // body{email,password}

	//favorites
	router.GET("/api/favorites", main.GET_Favorites)              //all favorites of CURRENT USER (auth)
	router.POST("/api/favorites/:id", main.POST_Favorite)         //add to fav by movieId of CURRENT USER (auth)
	router.DELETE("/api/favorites/:id", main.DELETE_Favorite)     //delete from favorites of CURRENT USER (auth) by movie id
	router.DELETE("/api/favorites/clear/", main.DELETE_Favorites) //clears all CURRENT USERS favorites

	router.Run(":8080")
}
