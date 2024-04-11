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
	// router.GET("/forgot", main.GET_Forgot)
	// router.POST("/api/forgot", main.POST_Forgot)                // not finished

	//HTML
	router.GET("/", main.GET_HTML_Index)                 // main page
	router.GET("/reg", main.GET_HTML_Reg)                // html of reg
	router.GET("/login", main.GET_HTML_Login)            // html of login
	router.GET("/create/project", main.GET_HTML_Project) // html of creating project

	//GET
	router.GET("/api/profile", main.GET_Profile)      // the user gets HIS profile
	router.GET("/api/projects", main.GET_Projects)    // get all projects, can limit projects by adding query ?limit=<number>
	router.GET("/api/projects/:id", main.GET_Project) // get project by id

	router.GET("/api/projects/season/:id", main.GET_Season)
	router.GET("/api/projects/episode/:id", main.GET_Episode) // get project by id

	//POST
	router.POST("/api/reg", main.POST_Reg)                // registration
	router.POST("/api/login", main.POST_Login)            // login registered user
	router.POST("/api/create/project", main.POST_Project) //create project only by admin role

	//DELETE
	router.DELETE("/api/delete/project/:id", main.DELETE_Project) // deletes project by id (admin)

	//PUT
	router.PUT("/api/update/project/:id", main.PUT_Project) // updates project by id (admin)

	router.Run(":8080")

}
