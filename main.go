package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/fgunawan1995/lemonilo/config"
	cachedal "github.com/fgunawan1995/lemonilo/dal/cache"
	dbdal "github.com/fgunawan1995/lemonilo/dal/db"
	"github.com/fgunawan1995/lemonilo/handler"
	"github.com/fgunawan1995/lemonilo/middleware"
	"github.com/fgunawan1995/lemonilo/resources"
	"github.com/fgunawan1995/lemonilo/usecase"
)

func main() {
	// Init resources
	cfg := config.GetConfig()
	db, err := resources.ConnectDb(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	redis, err := resources.ConnectRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	// Init layers
	cacheDAL := cachedal.New(redis)
	dbDAL := dbdal.New(db)
	usecaseLayer := usecase.New(cfg, dbDAL, cacheDAL)
	middlewareLayer := middleware.New(cacheDAL)
	handlerLayer := handler.New(usecaseLayer)

	// Init routes
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlerLayer.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/user", handlerLayer.InsertUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlerLayer.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", handlerLayer.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/login", handlerLayer.Login).Methods(http.MethodPost)
	r.HandleFunc("/test", middlewareLayer.ValidateJWT(handlerLayer.TestToken)).Methods(http.MethodGet)

	// Start server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(r)
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
