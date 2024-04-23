package mapping

import (
	"project/internal/database/user"

	"github.com/gin-gonic/gin"
)

func TrimUser(user user.User) gin.H {
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
	if user.Role != "" {
		userJson["role"] = user.Role
	}

	return userJson
}
func TrimUserJson(userVar *user.User) user.UserJson {
	var userJson user.UserJson

	if userVar.Id != 0 {
		userJson.Id = userVar.Id
	}
	if userVar.Email != "" {
		userJson.Email = userVar.Email
	}
	if userVar.Name.Valid {
		userJson.Name = userVar.Name.String
	}
	if userVar.Phone.Valid {
		userJson.Phone = userVar.Phone.String
	}

	if userVar.DateOfBirth.Valid {
		userJson.DateOfBirth = userVar.DateOfBirth.Time.Format("2006-01-02")
	}

	if userVar.Role != "" {
		userJson.Role = userVar.Role
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
