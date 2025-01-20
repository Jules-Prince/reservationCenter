package models

type Availability struct {
    ID        string `json:"id"`
    Date      string `json:"date"`
    StartTime string `json:"startTime"`
    EndTime   string `json:"endTime"`
    IsBooked  bool   `json:"isBooked"`
}