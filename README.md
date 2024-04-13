to run -> docker-compose up --build


creditnails for admin -> 

big@example.com
Aa12345678#



.
├── cmd
│   └── web
│       └── main.go
├── docker-compose.yml
├── dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── database
│   │   ├── dataset
│   │   │   ├── data.go
│   │   │   └── table.go
│   │   ├── episode
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── favorites
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   └── model.go
│   │   ├── models.go
│   │   ├── movie
│   │   │   ├── check.go
│   │   │   ├── create.go
│   │   │   ├── delete.go
│   │   │   ├── get.go
│   │   │   ├── model.go
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
│   │       └── model.go
│   ├── init
│   │   └── init.go
│   ├── routes
│   │   ├── episode.go
│   │   ├── favorites.go
│   │   ├── index.go
│   │   ├── init.go
│   │   ├── login.go
│   │   ├── middleware.go
│   │   ├── movie.go
│   │   ├── profile.go
│   │   ├── reg.go
│   │   ├── restore.go
│   │   ├── season.go
│   │   └── trends.go
│   ├── start
│   │   └── start.go
│   └── utils
│       ├── check.go
│       ├── code.go
│       ├── hashPassword.go
│       ├── jwt.go
│       └── mapping
│           └── mapping.go
├── README.md
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

23 directories, 68 files