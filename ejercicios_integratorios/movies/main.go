package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/ejercicios_integratorios/movies/controllers"
	"github.com/labora/labora-golang/ejercicios_integratorios/movies/services"
	"github.com/labora/labora-golang/ejercicios_integratorios/movies/util"
	"github.com/rs/cors"
)

func main() {
	// Load env data
	err := util.LoadEnv()
	if err != nil {
		return
	}

	envData := util.EnvData

	// Init server
	services.Init()
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movie", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.GetMovies).Methods("GET")

	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://" + envData.DbData.Host + ":" + envData.DbData.Port,
		},
		AllowedMethods: []string{
			"GET", "POST",
		},
	})
	handler := corsOptions.Handler(router)

	port := ":" + envData.ServerData.Port
	if err = initServer(port, handler); err != nil {
		return
	}
}

func initServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("🚀 Server started on port http://localhost%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("Error initializing the server. Error: %v\n", err)
	}

	return nil
}
