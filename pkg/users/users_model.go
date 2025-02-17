/*
Package Name: users
File Name: users_controller.go
Abstract: A representation of a user in the database.
*/
package users

import (
	"time"
)

// ======== TYPES ========

// InternalUser is a struct that represents a user, and it contains its password.
// As its own name suggests, this type should only be used internally.
type InternalUser struct {
	ID        int32
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

// PublicUser is basically a user that will be returned by the api. As its own
// name says, it should be used for returning user data publicly.
type PublicUser struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// ======== PUBLIC METHODS ========

// Converts an internal user to a public user.
func (self InternalUser) ToPublic() PublicUser {
	return PublicUser{
		ID:        self.ID,
		Username:  self.Username,
		Email:     self.Email,
		CreatedAt: self.CreatedAt,
	}
}

// Creates a new instance of an internal user from data.
func InternalUserFromData(values []interface{}) InternalUser {
	return InternalUser{
		ID:        values[0].(int32),
		Username:  values[1].(string),
		Email:     values[2].(string),
		Password:  values[3].(string),
		CreatedAt: values[4].(time.Time),
	}
}
