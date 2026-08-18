package models

import "time"

type Booking struct {
    ID      int       `json:"id"`
    UserID  int       `json:"user_id"`
    PlaceID int       `json:"place_id"`
    From    time.Time `json:"from"`
    To      time.Time `json:"to"`
}

type BookingsResponse struct {
    Bookings []Booking `json:"bookings"`
}
