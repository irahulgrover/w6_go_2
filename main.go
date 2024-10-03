package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Trip struct {
	ID          int    `json:"id"`
	Destination string `json:"destination"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

var trips []Trip

func createTrip(w http.ResponseWriter, r *http.Request) {
	var newTrip Trip
	json.NewDecoder(r.Body).Decode(&newTrip)
	newTrip.ID = len(trips) + 1
	trips = append(trips, newTrip)
	json.NewEncoder(w).Encode(newTrip)
}

func getTrips(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(trips)
}
func getTripByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, trip := range trips {
		if trip.ID == id {
			json.NewEncoder(w).Encode(trip)
			return
		}
	}
	http.Error(w, "Trip not found", http.StatusNotFound)
}
func updateTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, trip := range trips {
		if trip.ID == id {
			var updatedTrip Trip
			json.NewDecoder(r.Body).Decode(&updatedTrip)
			updatedTrip.ID = id
			trips[index] = updatedTrip
			json.NewEncoder(w).Encode(updatedTrip)
			return
		}
	}
	http.Error(w, "Trip not found", http.StatusNotFound)
}
func deleteTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, trip := range trips {
		if trip.ID == id {
			trips = append(trips[:index], trips[index+1:]...)
			json.NewEncoder(w).Encode(trips)
			return
		}
	}
	http.Error(w, "Trip not found", http.StatusNotFound)
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/trips", createTrip).Methods("POST")
	router.HandleFunc("/trips", getTrips).Methods("GET")
	router.HandleFunc("/trips/{id}", getTripByID).Methods("GET")
	router.HandleFunc("/trips/{id}", updateTrip).Methods("PUT")
	router.HandleFunc("/trips/{id}", deleteTrip).Methods("DELETE")

	// Similar routes for /locations and /activities
	fmt.Println("Server started at:8000")
	http.ListenAndServe(":8000", router)
}
