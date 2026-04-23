package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mijaelr2003/staycore/backend/core/models"
	"github.com/mijaelr2003/staycore/backend/database"
)

func GetBookings(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking
	database.DB.Find(&bookings)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    bookings,
	})
}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking

	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    booking,
	})
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if booking.RoomID == (uuid.UUID{}) || booking.GuestID == (uuid.UUID{}) {
		http.Error(w, "room_id and guest_id are required", http.StatusBadRequest)
		return
	}

	if booking.CheckInDate.IsZero() || booking.CheckOutDate.IsZero() {
		http.Error(w, "check_in_date and check_out_date are required", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    booking,
	})
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking

	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	database.DB.Save(&booking)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    booking,
	})
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking

	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	database.DB.Model(&booking).Update("status", "cancelled")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Booking cancelled",
	})
}

func CheckInBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking

	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	if booking.Status != "confirmed" && booking.Status != "pending" {
		http.Error(w, "Booking is not in a valid state for check-in", http.StatusBadRequest)
		return
	}

	checkin := models.CheckIn{
		BookingID: booking.ID,
	}

	if err := database.DB.Create(&checkin).Error; err != nil {
		http.Error(w, "Error creating check-in", http.StatusInternalServerError)
		return
	}

	database.DB.Model(&booking).Update("status", "checked_in")
	database.DB.Model(&models.Room{}).Where("id = ?", booking.RoomID).Update("status", "occupied")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    checkin,
	})
}

func CheckOutBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking

	if err := database.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	if booking.Status != "checked_in" {
		http.Error(w, "Booking is not checked in", http.StatusBadRequest)
		return
	}

	checkout := models.CheckOut{
		BookingID: booking.ID,
	}

	if err := database.DB.Create(&checkout).Error; err != nil {
		http.Error(w, "Error creating check-out", http.StatusInternalServerError)
		return
	}

	database.DB.Model(&booking).Update("status", "checked_out")
	database.DB.Model(&models.Room{}).Where("id = ?", booking.RoomID).Update("status", "available")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    checkout,
	})
}