package main

import(
    "time"
)

type Event struct {
  Id          int         `json:"id"`
  EventName   string      `json:"event_name" validate:"required"`
  RoomName    string      `json:"room_name"`
  Description string      `json:"description"`
  Items       string      `json:"items" validate:"required"`
  Major       int         `json:"major"`
  CreatedAt   time.Time   `json:"created_at"`
  DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}

type User struct {
  Id            int     `json:"id"`
  UserName      string  `json:"user_name" validate:"required"`
  Profile       string  `json:"profile"`
  Items         string  `json:"items" validate:"required"`
  Major         int     `json:"major"`
  Image         string  `json:"image"`
  ImageHeader 	string  `json:"image_header"`
  CreatedAt   time.Time `json:"created_at"`
  DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}
