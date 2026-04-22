package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mijaelr2003/staycore/backend/core/models"
	"github.com/mijaelr2003/staycore/backend/database"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	propertyID := mux.Vars(r)["id"]
	var rooms []models.Room

	database.DB.Where("property_id = ? AND active = ?", propertyID, true).Find(&rooms)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    rooms,
	})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	propertyID := vars["id"]
	roomID := vars["rid"]
	var room models.Room

	if err := database.DB.Where("id = ? AND property_id = ? AND active = ?", roomID, propertyID, true).First(&room).Error; err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    room,
	})
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	propertyID := mux.Vars(r)["id"]
	var room models.Room

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if room.Number == "" || room.Type == "" {
		http.Error(w, "number and type are required", http.StatusBadRequest)
		return
	}

	// Verificar que la property existe
	var property models.Property
	if err := database.DB.Where("id = ? AND active = ?", propertyID, true).First(&property).Error; err != nil {
		http.Error(w, "Property not found", http.StatusNotFound)
		return
	}

	// Parsear el UUID de la property
	room.PropertyID = property.ID

	if err := database.DB.Create(&room).Error; err != nil {
		http.Error(w, "Error creating room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    room,
	})
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	propertyID := vars["id"]
	roomID := vars["rid"]
	var room models.Room

	if err := database.DB.Where("id = ? AND property_id = ?", roomID, propertyID).First(&room).Error; err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	database.DB.Save(&room)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    room,
	})
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	propertyID := vars["id"]
	roomID := vars["rid"]
	var room models.Room

	if err := database.DB.Where("id = ? AND property_id = ?", roomID, propertyID).First(&room).Error; err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	database.DB.Model(&room).Update("active", false)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message": "Room deleted",
	})
}