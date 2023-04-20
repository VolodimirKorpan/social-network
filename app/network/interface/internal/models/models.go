package models

import "time"

type User struct {
	ID       string
	Username string
	Password string
	Avatar   string
	Bio      string
	Friends  []*Friendship
}

type Friendship struct {
	RequesterID string 
	RequesteeID string 
	Status      string
	RequestedAt time.Time
	ResolvedAt  time.Time
	Requester   User 
	Requestee   User 
}
