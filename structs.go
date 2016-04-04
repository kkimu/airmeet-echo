package main

import(
    "time"
)

type Event struct {
  Id          int     `json:"id"`
  EventName   string    `json:"event_name" validate:"required"`
  RoomName    string    `json:"room_name"`
  Description string    `json:"description"`
  Items       string    `json:"items" validate:"required"`
  Major       int     `json:"major"`
  Active      bool      `json:"active" validate:"required"`
  CreatedAt   time.Time `json:"created_at"`
  DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}

type User struct {
  Id        int
  UserName  string
  Profile   string
}
