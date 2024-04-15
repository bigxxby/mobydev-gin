package movies

import (
	"database/sql"
	"log"
	"project/internal/database/age"
	"project/internal/database/categories"
	"project/internal/database/genres"
	"project/internal/utils"

	"github.com/gin-gonic/gin"
)

func (m *MoviesRoute) GET_Movies_MAIN(c *gin.Context) {
	limit := c.Query("limit")
	if limit == "" {
		movies, err := m.DB.MovieRepository.GetMovies(0)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "no movies found",
				})
				return
			}
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		agesAll, err := m.DB.AgeRepository.GetAllAgeCategories()
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "no ages found",
				})
				return
			}
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		categoriesAll, err := m.DB.CategoriesRepository.GetCategories()
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "no categories found",
				})
				return
			}
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}

		genresAll, err := m.DB.GenreRepository.GetAllGenres()
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{
					"message": "no genres found",
				})
				return
			}
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		categoryMap := make(map[int]categories.Category)
		for _, cat := range categoriesAll {
			categoryMap[cat.ID] = cat
		}

		ageCategoryMap := make(map[int]age.AgeCategory)
		for _, ageCat := range agesAll {
			ageCategoryMap[ageCat.ID] = ageCat
		}

		genreMap := make(map[int]genres.Genre)
		for _, gen := range genresAll {
			genreMap[gen.ID] = gen
		}

		for i, movie := range movies {
			movies[i].Categories = []categories.Category{categoryMap[movie.CategoryId]}
			movies[i].AgeCategories = []age.AgeCategory{ageCategoryMap[movie.AgeCategoryId]}
			movies[i].Genres = []genres.Genre{genreMap[movie.GenreId]}
		}

		c.JSON(200, gin.H{
			"movies": movies,
		})
		return
	}
	valid, num := utils.IsValidNum(limit)
	if !valid {
		log.Println("number not valid")
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	movies, err := m.DB.MovieRepository.GetMovies(num)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "no movies found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	agesAll, err := m.DB.AgeRepository.GetAllAgeCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "no ages found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	categoriesAll, err := m.DB.CategoriesRepository.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "no categories found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	genresAll, err := m.DB.GenreRepository.GetAllGenres()
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{
				"message": "no genres found",
			})
			return
		}
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	categoryMap := make(map[int]categories.Category)
	for _, cat := range categoriesAll {
		categoryMap[cat.ID] = cat
	}

	ageCategoryMap := make(map[int]age.AgeCategory)
	for _, ageCat := range agesAll {
		ageCategoryMap[ageCat.ID] = ageCat
	}

	genreMap := make(map[int]genres.Genre)
	for _, gen := range genresAll {
		genreMap[gen.ID] = gen
	}

	for i, movie := range movies {
		movies[i].Categories = []categories.Category{categoryMap[movie.CategoryId]}
		movies[i].AgeCategories = []age.AgeCategory{ageCategoryMap[movie.AgeCategoryId]}
		movies[i].Genres = []genres.Genre{genreMap[movie.GenreId]}
	}

	c.JSON(200, gin.H{
		"movies": movies,
	})

}
