package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mijaelr2003/staycore/backend/core/models"
	"github.com/mijaelr2003/staycore/backend/database"
)

func GetGuests(w http.ResponseWriter, r *http.Request) {
	var guests []models.Guest
	database.DB.Find(&guests)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    guests,
	})
}

func CreateGuest(w http.ResponseWriter, r *http.Request) {
	var guest models.Guest

	if err := json.NewDecoder(r.Body).Decode(&guest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if guest.FirstName == "" || guest.LastName == "" || guest.DocumentNumber == "" {
		http.Error(w, "first_name, last_name and document_number are required", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&guest).Error; err != nil {
		http.Error(w, "Error creating guest", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    guest,
	})
}