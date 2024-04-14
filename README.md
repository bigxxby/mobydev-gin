to run -> docker-compose up --build


creditnails for admin -> 

big@example.com
Aa12345678#


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
│   │   ├── model.go
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
│   └── routes
│       ├── api
│       │   ├── auth
│       │   │   ├── model.go
│       │   │   ├── restore.go
│       │   │   ├── signIn.go
│       │   │   └── signUp.go
│       │   ├── episodes
│       │   │   ├── GET.go
│       │   │   └── model.go
│       │   ├── favorites
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   └── POST.go
│       │   ├── movies
│       │   │   ├── DELETE.go
│       │   │   ├── GET.go
│       │   │   ├── model.go
│       │   │   ├── POST.go
│       │   │   └── PUT.go
│       │   ├── seasons
│       │   │   ├── GET.go
│       │   │   └── model.go
│       │   ├── trends
│       │   │   ├── GET.go
│       │   │   └── model.go
│       │   └── users
│       │       ├── GET.go
│       │       └── model.go
│       ├── htmlHandlers.go
│       └── routeManager.go
├── README.md
├── test
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