// Golang Standards Project Layout no GitHub.

package main

import (
	"10-apis/configs"
	_ "10-apis/docs"
	"10-apis/internal/entity"
	"10-apis/internal/infra/db"
	"10-apis/internal/infra/webserver/handlers"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Expert API Example
// @version 1.0
// @description This is a sample server for Go Expert API.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 0JqFP@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config)

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := db.NewProductDB(database)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := db.NewUserDB(database)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
	// r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate-token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler())

	http.ListenAndServe(":8000", r)
}

// Ver Go Fiber, Go Echo, Gin
// frameworks web server

// Ver Buffalo
// framework Go completo

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf(time.Now().String(), r.Method, r.RequestURI)
// 		next.ServeHTTP(w, r)
// 	})
// }
