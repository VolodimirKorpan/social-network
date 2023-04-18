package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username   string `gorm:"not null,unique_index"`
	Password   string `gorm:"not null"`
	Avatar     string
	Bio        string
	Followers  []Follow `gorm:"foreignkey:FollowingID"`
	Followings []Follow `gorm:"foreignkey:FollowerID"`
}

type Follow struct {
	gorm.Model
	Follower    User
	FollowerID  string `gorm:"primaryKey" sql:"type:uuid not null"`
	Following   User
	FollowingID string `gorm:"primary_key" sql:"type:uuid not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
