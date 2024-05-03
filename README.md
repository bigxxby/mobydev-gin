# ozinshe-test-app

This project is a movie streaming platform that allows users to browse, search, and watch movies. It also provides features for administrators to manage movies, categories, and user profiles.

## Running the Application

To run the application, execute the following commands in your terminal:

```bash
docker-compose build
docker-compose up

    OR

cd .
go run ./server
```

## Entry point

```
http://localhost:8080/docs/index.html
```

## Admin Credentials

Use the following credentials to log in as an admin:

- **Email:** big@example.com
- **Password:** Aa12345678#

## Features in Progress

### Frontend

- **Status:** NOT DONE

### Unit-tests

- **Status:** NOT DONE

## General Notes:

You must include a valid authentication token in the request headers if its a route that require admin privilege or user auth.
Admin routes are indicated where only administrators have access. Admin token saved in local storage after login.

```
.
├── checklist
├── docker-compose.yml
├── dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
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
│   │   │   ├── drop.go
│   │   │   ├── init.go
│   │   │   └── tables.go
│   │   ├── episode
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   └── update.go
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
│   │   │   ├── getOptions.go
│   │   │   ├── model.go
│   │   │   ├── search.go
│   │   │   └── update.go
│   │   ├── posters
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── season
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
│   │   │   └── update.go
│   │   └── user
│   │       ├── check.go
│   │       ├── create.go
│   │       ├── delete.go
│   │       ├── get.go
│   │       ├── model.go
│   │       └── update.go
│   ├── init
│   │   └── initDb.go
│   └── utils
│       ├── auth.go
│       ├── codes.go
│       ├── encryption.go
│       ├── mapping
│       │   └── mapping.go
│       ├── random.go
│       ├── smtp.go
│       └── validation.go
├── pkg
│   ├── middleware
│   │   └── auth.go
│   └── routes
│       ├── api
│       │   ├── age
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   ├── auth
│       │   │   ├── changePassword.go
│       │   │   ├── model.go
│       │   │   ├── old
│       │   │   │   └── restorePassword.go
│       │   │   ├── signIn.go
│       │   │   └── signUp.go
│       │   ├── categories
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   ├── episodes
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   ├── favorites
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   └── POST.go
│       │   ├── genres
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   ├── movies
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   ├── PUT_ageCategory.go
│       │   │   ├── PUT_categoires.go
│       │   │   ├── PUT_data.go
│       │   │   ├── PUT_genres.go
│       │   │   └── query.go
│       │   ├── posters
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   └── POST.go
│       │   ├── seasons
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   └── users
│       │       ├── GET.go
│       │       ├── model.go
│       │       └── PUT.go
│       ├── routeManager.go
│       └── swaggerModels.go
├── README.md
└── server
    └── main.go

32 directories, 132 files
```
