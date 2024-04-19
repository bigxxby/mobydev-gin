# ozinshe-test-app

This project is a movie streaming platform that allows users to browse, search, and watch movies. It also provides features for administrators to manage movies, categories, and user profiles.

## Running the Application

To run the application, execute the following commands in your terminal:

```bash
docker-compose build
docker-compose up
```

## Admin Credentials

Use the following credentials to log in as an admin:

- **Email:** big@example.com
- **Password:** Aa12345678#

## Features in Progress

### Reset Password

- **Status:** NOT DONE

### Search for Genre Movie

- **Status:** NOT DONE

### Swagger docs

- **Status:** NOT DONE

### Frontend

- **Status:** NOT DONE

### Unit-tests

- **Status:** NOT DONE


## General Notes:
You must include a valid authentication token in the request headers if its a route that require admin privilege or user auth.
Admin routes are indicated where only administrators have access. Admin token saved in local storage after login.


## API Reference

### Movies

#### Retrieve All Movies

- `GET /api/movies`: Retrieve a list of movies. Optional query parameter `limit` can be used to limit the number of movies returned.

#### Retrieve Movie by ID

- `GET /api/movies/:id`: Retrieve details of a movie by its ID.

#### Add New Movie

- `POST /api/movies`: Add a new movie. **Admin only**.

#### Delete Movie by ID

- `DELETE /api/movies/:id`: Delete a movie by its ID. **Admin only**.

#### Update Movie by ID

- `PUT /api/movies/:id`: Update details of a movie by its ID. **Admin only**.

#### Retrieve Main Movies

- `GET /api/movies/main`: Retrieve movies with categories, age category, and genre. Optional query parameter `limit` can be used.

#### Search Movies

- `GET /api/movies/search`: Search for movies based on a query. Use `query` parameter to specify the search query. 

### Seasons

#### Retrieve Seasons by Movie ID

- `GET /api/seasons/:id`: Retrieve all seasons of a specific movie by its ID.

### Episodes

#### Retrieve Episodes by Season ID

- `GET /api/seasons/:id`: Retrieve all episodes of a specific season by its ID.

### Profile

#### Retrieve Profile

- `GET /api/profile`: Retrieve profile details of the currently authenticated user. **User auth**.

#### Update Profile

- `PUT /api/profile`: Update profile details of the currently authenticated user. Fields that can be updated include `dob`, `name`, and `phone`. **User auth**.

### Trends

#### Retrieve Trend by ID

- `GET /api/trends/:id`: Retrieve details of a trend by its ID.

#### Retrieve All Trends

- `GET /api/trends`: Retrieve all current trends.

### Authentication

#### Register New User

- `POST /api/signUp`: Register a new user. Required fields: `email`, `password`, `role`.

#### Sign In User

- `POST /api/signIn`: Authenticate and sign in a user. Required fields: `email`, `password`. 

### Favorites

#### Retrieve Favorites

- `GET /api/favorites`: Retrieve favorite movies of the currently authenticated user.  **User auth**.

#### Add Favorite

- `POST /api/favorites/:id`: Add a movie to favorites for the currently authenticated user. **User auth**.

#### Remove Favorite

- `DELETE /api/favorites/:id`: Remove a movie from favorites of the currently authenticated user. **User auth**.

#### Clear Favorites

- `DELETE /api/favorites/clear`: Remove all movies from favorites of the currently authenticated user. **User auth**.

### Categories

#### Retrieve All Categories

- `GET /api/categories`: Retrieve all movie categories.

#### Retrieve Category by ID

- `GET /api/categories/:id`: Retrieve a movie category by its ID.

#### Add Category

- `POST /api/categories`: Add a new movie category. **Admin only**.

#### Update Category

- `PUT /api/categories/:id`: Update a movie category by its ID. **Admin only**.

#### Delete Category

- `DELETE /api/categories/:id`: Delete a movie category by its ID. **Admin only**.

### Genres

#### Retrieve All Genres

- `GET /api/genres`: Retrieve all movie genres.

#### Retrieve Genre by ID

- `GET /api/genres/:id`: Retrieve a movie genre by its ID.

#### Add Genre

- `POST /api/genres`: Add a new movie genre. **Admin only**.

#### Update Genre

- `PUT /api/genres/:id`: Update a movie genre by its ID. **Admin only**.

#### Delete Genre

- `DELETE /api/genres/:id`: Delete a movie genre by its ID. **Admin only**.

### Age Categories

#### Retrieve All Age Categories

- `GET /api/ageCategories`: Retrieve all age categories.

#### Retrieve Age Category by ID

- `GET /api/ageCategories/:id`: Retrieve an age category by its ID. 

#### Add Age Category

- `POST /api/ageCategories`: Add a new age category. **Admin only**.

#### Update Age Category

- `PUT /api/ageCategories/:id`: Update an age category by its ID. **Admin only**.

#### Delete Age Category

- `DELETE /api/ageCategories/:id`: Delete an age category by its ID. **Admin only**.





```
.
├── app
│   └── main.go
├── docker-compose.yml
├── dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── database
│   │   ├── age
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   └── update.go
│   │   ├── categories
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   └── update.go
│   │   ├── datasets
│   │   │   ├── createTables.go
│   │   │   ├── createTestData.go
│   │   │   ├── data.go
│   │   │   ├── init.go
│   │   │   └── tables.go
│   │   ├── episode
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── favorites
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── genres
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   └── update.go
│   │   ├── modelRepositories.go
│   │   ├── movie
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   ├── search.go
│   │   │   └── update.go
│   │   ├── season
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── trend
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   └── user
│   │       ├── check.go
│   │       ├── create.go
│   │       ├── get.go
│   │       ├── model.go
│   │       └── update.go
│   ├── init
│   │   └── initDb.go
│   ├── start
│   │   └── start.go
│   └── utils
│       ├── auth.go
│       ├── code.go
│       ├── encryption.go
│       ├── mapping
│       │   └── mapping.go
│       └── validation.go
├── pkg
│   ├── middleware
│   │   └── auth.go
│   ├── routes
│   │   ├── api
│   │   │   ├── age
│   │   │   │   ├── DELETE.go
│   │   │   │   ├── GET.go
│   │   │   │   ├── model.go
│   │   │   │   ├── POST.go
│   │   │   │   └── UPDATE.go
│   │   │   ├── auth
│   │   │   │   ├── model.go
│   │   │   │   ├── restore.go
│   │   │   │   ├── signIn.go
│   │   │   │   └── signUp.go
│   │   │   ├── categories
│   │   │   │   ├── DELETE.go
│   │   │   │   ├── GET.go
│   │   │   │   ├── model.go
│   │   │   │   ├── POST.go
│   │   │   │   └── PUT.go
│   │   │   ├── episodes
│   │   │   │   ├── GET.go
│   │   │   │   └── model.go
│   │   │   ├── favorites
│   │   │   │   ├── DELETE.go
│   │   │   │   ├── GET.go
│   │   │   │   ├── model.go
│   │   │   │   └── POST.go
│   │   │   ├── genres
│   │   │   │   ├── DELETE.go
│   │   │   │   ├── GET.go
│   │   │   │   ├── model.go
│   │   │   │   ├── POST.go
│   │   │   │   └── PUT.go
│   │   │   ├── movies
│   │   │   │   ├── DELETE.go
│   │   │   │   ├── GET.go
│   │   │   │   ├── GET_MAIN.go
│   │   │   │   ├── model.go
│   │   │   │   ├── POST.go
│   │   │   │   ├── PUT.go
│   │   │   │   └── search.go
│   │   │   ├── seasons
│   │   │   │   ├── GET.go
│   │   │   │   └── model.go
│   │   │   ├── trends
│   │   │   │   ├── GET.go
│   │   │   │   └── model.go
│   │   │   └── users
│   │   │       ├── GET.go
│   │   │       ├── model.go
│   │   │       └── PUT.go
│   │   ├── htmlHandlers.go
│   │   └── routeManager.go
│   └── search
│       └── search.go
├── README.md
├── test
│   └── get_test.go
├── test.txt
└── ui
    ├── static
    │   ├── css
    │   │   ├── index.css
    │   │   ├── login.css
    │   │   └── reg.css
    │   ├── images
    │   │   ├── back1-dark.jpg
    │   │   ├── back1.jpg
    │   │   ├── back2-dark.jpg
    │   │   ├── back2.jpg
    │   │   ├── default.jpg
    │   │   └── icons
    │   │       ├── logo.png
    │   │       └── union.png
    │   └── script
    │       ├── index.js
    │       ├── login.js
    │       └── reg.js
    └── templates
        ├── home.html
        ├── index.html
        ├── login.html
        ├── project_create.html
        ├── reg.html
        └── restore.html

40 directories, 124 files
```
