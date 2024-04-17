package test

import (
	"net/http"
	"net/http/httptest"
	"project/internal/database/datasets"
	"project/pkg/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var main *routes.Manager

func Setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	main, err = routes.Init()
	if err != nil {
		panic(err)
	}

	err = datasets.InitDatasets(main.DB)
	if err != nil {
		panic(err)
	}
}

func TestGetMovies(t *testing.T) {
	Setup()

	router := gin.Default()
	router.GET("/api/movies/", main.MoviesRoute.GET_Movies)

	req, err := http.NewRequest("GET", "/api/movies/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetMovieByID(t *testing.T) {
	Setup()

	router := gin.Default()
	router.GET("/api/movies/:id", main.MoviesRoute.GET_Movie)

	req, err := http.NewRequest("GET", "/api/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}
func TestGetEpisodeByID(t *testing.T) {
	Setup()

	router := gin.Default()
	episodes := router.Group("/api/episodes")
	{
		episodes.GET("/:id", main.EpisodesRoute.GET_Episode)
	}

	req, err := http.NewRequest("GET", "/api/episodes/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetSeasonByID(t *testing.T) {
	Setup()

	router := gin.Default()
	seasons := router.Group("/api/seasons")
	{
		seasons.GET("/:id", main.SeasonsRoute.GET_Season)
	}

	req, err := http.NewRequest("GET", "/api/seasons/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetTrendByID(t *testing.T) {
	Setup()

	router := gin.Default()
	trends := router.Group("/api/trends")
	{
		trends.GET("/:id", main.TrendsRoute.GET_Trend)
	}

	req, err := http.NewRequest("GET", "/api/trends/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetTrends(t *testing.T) {
	Setup()

	router := gin.Default()
	trends := router.Group("/api/trends")
	{
		trends.GET("/", main.TrendsRoute.GET_Trends)
	}

	req, err := http.NewRequest("GET", "/api/trends/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}

func TestGetCategories(t *testing.T) {
	Setup()

	router := gin.Default()
	categories := router.Group("/api/categories")
	{
		categories.GET("/", main.CategoriesRoute.GET_Categories)
	}

	req, err := http.NewRequest("GET", "/api/categories/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status OK")
}
