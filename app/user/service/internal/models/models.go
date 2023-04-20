package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `sql:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Avatar   string
	Bio      string
	Friends  []*Friendship `gorm:"foreignKey:RequesterID;references:ID"`
}

type Friendship struct {
	RequesterID string `sql:"type:uuid" gorm:"size:36"`
	RequesteeID string `sql:"type:uuid" gorm:"size:36"`
	Status      string
	RequestedAt time.Time
	ResolvedAt  time.Time
	Requester   User `gorm:"foreignkey:RequesterID"`
	Requestee   User `gorm:"foreignkey:RequesteeID"`
}

type FriendshipStatus string

const (
	FriendshipStatusPending  string = "pending"
	FriendshipStatusAccepted string = "accepted"
	FriendshipStatusRejected string = "rejected"
)


// BeforeCreate is a callback function that is called before creating a new user record.
// It generates a new UUID and assigns it to the user's ID field.
// The function takes a *gorm.DB object as the input parameter and returns an error object.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID string using the uuid package
	u.ID = uuid.NewString()
	return
}
