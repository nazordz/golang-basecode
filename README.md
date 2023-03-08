# Golang Basecode

## Dependecies
1. [gin](github.com/gin-gonic/gin) - Routing, middleware, & form validation
2. [gorm](gorm.io/gorm) - ORM & migration
3. [logrus](github.com/sirupsen/logrus) - Logging
4. [godotenv](github.com/joho/godotenv) - Environment variable
5. [wire](github.com/google/wire) - Dependency Injection
6. [jwt](github.com/golang-jwt/jwt) - JSON Web Token for authentication
7. [air](github.com/cosmtrek/air) - Auto refresh

## Folder structures
```bash
├── Dockerfile
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
├── injector.go
├── internal
│   ├── config
│   ├── controllers
│   │   ├── authentication_controller.go
│   │   └── user_controller.go
│   ├── models
│   │   └── user.go
│   ├── repositories
│   │   └── user_repository.go
│   ├── routers
│   │   ├── authentication_router.go
│   │   ├── routes.go
│   │   └── user_router.go
│   ├── services
│   │   └── user_service.go
│   ├── test
│   └── utils
│       ├── token.go
│       └── utils.go
├── main.go
├── pkg
│   ├── database.go
│   ├── dotenv.go
│   └── middlewares
│       └── middleware.go
├── tmp
│   ├── build-errors.log
│   └── main
└── wire_gen.go
```