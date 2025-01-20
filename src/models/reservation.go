package models

type Reservation struct {
    ID           string `json:"id"`
    AvailabilityID string `json:"availabilityId"`
    UserID       string `json:"userId"`
    CreatedAt    string `json:"createdAt"`
}
