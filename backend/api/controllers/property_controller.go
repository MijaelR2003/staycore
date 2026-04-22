package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mijaelr2003/staycore/backend/core/models"
	"github.com/mijaelr2003/staycore/backend/database"
)

func GetProperties(w http.ResponseWriter, r *http.Request) {
	var properties []models.Property
	database.DB.Where("active = ?", true).Find(&properties)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    properties,
	})
}

func GetProperty(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var property models.Property

	if err := database.DB.Where("id = ? AND active = ?", id, true).First(&property).Error; err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    property,
	})
}

func CreateProperty(w http.ResponseWriter, r *http.Request) {
	var property models.Property

	if err := json.NewDecoder(r.Body).Decode(&property); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if property.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&property).Error; err != nil {
		http.Error(w, "Error creating property", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    property,
	})
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var property models.Property

	if err := database.DB.Where("id = ?", id).First(&property).Error; err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&property); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	database.DB.Save(&property)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    property,
	})
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var property models.Property

	if err := database.DB.Where("id = ?", id).First(&property).Error; err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Soft delete — solo marca como inactiva
	database.DB.Model(&property).Update("active", false)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Property deleted",
	})
}