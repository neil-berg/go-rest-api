package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/neil-berg/go-rest/handlers"
)

func main() {
	logger := log.New(os.Stdout, "recipe-api ", log.LstdFlags)

	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err)
	}

	handler := handlers.CreateHandler(logger)

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/recipes", handler.GetRecipes)
	getRouter.HandleFunc("/users", handler.GetUsers)

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/recipes", handler.AddRecipe)
	postRouter.Use(handler.ParseJSONRecipe)

	putRouter := router.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/recipes/{id:[\\w]+}", handler.UpdateRecipe)
	putRouter.Use(handler.ParseJSONRecipe)

	// CORS (for demo client on 9000)
	allowedOrigins := []string{"http://localhost:9000"}
	cors := ghandlers.CORS(ghandlers.AllowedOrigins(allowedOrigins))

	port := os.Getenv("SERVER_PORT")
	address := ":" + port

	s := http.Server{
		Addr:         address,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("The server is running on port", port)
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Printf("Recieved terminal signal [%s], gracefully shutting down... \n", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
