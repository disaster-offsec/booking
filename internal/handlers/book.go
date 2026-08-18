
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"booking/internal/storage"
)

func Book(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	
	placeID, err := strconv.Atoi(query.Get("place_id"))
	if err != nil || placeID <= 0 {
		http.Error(w, "invalid place_id", http.StatusBadRequest)
		return
	}
	
	userID, err := strconv.Atoi(query.Get("user_id"))
	if err != nil || userID <= 0 {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	
	from, err := time.Parse(time.RFC3339, query.Get("from"))
	if err != nil {
		http.Error(w, "invalid from", http.StatusBadRequest)
		return
	}
	
	to, err := time.Parse(time.RFC3339, query.Get("to"))
	if err != nil {
		http.Error(w, "invalid to", http.StatusBadRequest)
		return
	}
	
	if !to.After(from) {
		http.Error(w, "from must be before to", http.StatusBadRequest)
		return
	}

	var conflict bool
	storage.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM bookings 
			WHERE place_id=$1 AND time_from<$2 AND $3<time_to
		)`, placeID, to, from).Scan(&conflict)
	
	if conflict {
		http.Error(w, "Conflict", http.StatusConflict)
		return
	}
	
	var id int
	storage.DB.QueryRow(`
		INSERT INTO bookings (user_id, place_id, time_from, time_to)
		VALUES ($1,$2,$3,$4) RETURNING id`,
		userID, placeID, from, to).Scan(&id)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
