package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mijaelr2003/staycore/backend/api/controllers"
	"github.com/rs/cors"
)

func Start() {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": "staycore-backend",
		})
	}).Methods("GET")

	// API v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// Guests
	api.HandleFunc("/guests", controllers.GetGuests).Methods("GET")
	api.HandleFunc("/guests", controllers.CreateGuest).Methods("POST")

	// Properties
	api.HandleFunc("/properties", controllers.GetProperties).Methods("GET")
	api.HandleFunc("/properties", controllers.CreateProperty).Methods("POST")
	api.HandleFunc("/properties/{id}", controllers.GetProperty).Methods("GET")
	api.HandleFunc("/properties/{id}", controllers.UpdateProperty).Methods("PUT")
	api.HandleFunc("/properties/{id}", controllers.DeleteProperty).Methods("DELETE")

	// Rooms
	api.HandleFunc("/properties/{id}/rooms", controllers.GetRooms).Methods("GET")
	api.HandleFunc("/properties/{id}/rooms", controllers.CreateRoom).Methods("POST")
	api.HandleFunc("/properties/{id}/rooms/{rid}", controllers.GetRoom).Methods("GET")
	api.HandleFunc("/properties/{id}/rooms/{rid}", controllers.UpdateRoom).Methods("PUT")
	api.HandleFunc("/properties/{id}/rooms/{rid}", controllers.DeleteRoom).Methods("DELETE")

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	handler := c.Handler(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}