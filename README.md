# Golang Basecode

## Prerequisites
1. Golang v1.20.0
2. Postgresql > v12

## Dependecies
1. [gin](https://github.com/gin-gonic/gin) - Routing, middleware, & form validation
2. [gorm](https://gorm.io/gorm) - ORM & migration
3. [logrus](https://github.com/sirupsen/logrus) - Logging
4. [godotenv](https://github.com/joho/godotenv) - Environment variable
5. [wire](https://github.com/google/wire) - Dependency Injection
6. [jwt](https://github.com/golang-jwt/jwt) - JSON Web Token for authentication
7. [air](https://github.com/cosmtrek/air) - Auto refresh


## Getting Started
1. copy `.env.example` to `.env`
2. run `go install github.com/cosmtrek/air@latest` to install [air](https://github.com/cosmtrek/air)
3. run `air` in terminal to run in development

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