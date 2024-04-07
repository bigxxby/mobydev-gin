package mapping

import (
	"project/internal/database"

	"github.com/gin-gonic/gin"
)

func TrimUser(user database.User) gin.H {
	userJson := gin.H{}

	userJson["id"] = user.Id
	userJson["email"] = user.Email
	if user.Name.Valid {
		userJson["name"] = user.Name.String
	}
	if user.Phone.Valid {
		userJson["phone"] = user.Phone.String
	}
	if user.DateOfBirth.Valid {
		userJson["dot"] = user.DateOfBirth.Time.String()
	}

	return userJson
}

// func TrimMovies(movies []database.Movie) []database.MovieJson {
// 	var trimmedMovies []database.MovieJson

// 	for _, movie := range movies {
// 		trimmedMovie := database.MovieJson{
// 			Id:              movie.Id,
// 			Name:            movie.Name.String,
// 			Category:        movie.Category.String,
// 			ProjectType:     movie.ProjectType.String,
// 			Year:            int(movie.Year.Int32),
// 			AgeCategory:     movie.AgeCategory.String,
// 			DurationMinutes: int(movie.DurationMinutes.Int32),
// 			Keywords:        movie.Keywords.String,
// 			Description:     movie.Description.String,
// 			Director:        movie.Director.String,
// 			Producer:        movie.Producer.String,
// 		}

// 		trimmedMovies = append(trimmedMovies, trimmedMovie)
// 	}

// 	return trimmedMovies
// }
