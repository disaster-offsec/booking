package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"booking/internal/models"
	"booking/internal/storage"
)

func Booklist(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	uid := q.Get("user_id")
	pid := q.Get("place_id")

	if (uid == "" && pid == "") || (uid != "" && pid != "") {
		http.Error(w, "need exactly one param", http.StatusBadRequest)
		return
	}

	var rows *sql.Rows
	var err error

	if uid != "" {
		rows, err = storage.DB.Query(`
			SELECT id, user_id, place_id, time_from, time_to
			FROM bookings WHERE user_id=$1
			ORDER BY time_from, id`, uid)
	} else {
		rows, err = storage.DB.Query(`
			SELECT id, user_id, place_id, time_from, time_to
			FROM bookings WHERE place_id=$1
			ORDER BY time_from, id`, pid)
	}

	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	list := []models.Booking{}
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.UserID, &b.PlaceID, &b.From, &b.To); err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
		list = append(list, b)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.BookingsResponse{Bookings: list})
}
