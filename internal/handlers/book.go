
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
	
	placeID, _ := strconv.Atoi(query.Get("place_id"))
	userID, _ := strconv.Atoi(query.Get("user_id"))
	from, _ := time.Parse(time.RFC3339, query.Get("from"))
	to, _ := time.Parse(time.RFC3339, query.Get("to"))
	
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
