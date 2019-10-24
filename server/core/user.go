package core

import (
	"context"
	"time"
)

type (
	// User represents model for users
	User struct {
		ID           ID        `json:"id"`
		Username     string    `json:"username"`
		Email        string    `json:"email"`
		Password     []byte    `json:"-"`
		LastLogin    time.Time `json:"-"`
		Active       bool      `json:"active"`
		Admin        bool      `json:"admin"`
		RefreshToken string    `json:"-"`
	}

	// UserStore defines operations for working with persistent storage
	UserStore interface {
		FindByID(id ID) (*User, error)

		FindByUsername(username string) (*User, error)

		// IsExist checks if a user with username exists in database
		IsExists(username string) (bool, error)

		// List returns a Users list from database
		List() (result []User, err error)

		Insert(*User) (newID ID, err error)

		Update(*User) error
	}

	UserService interface {

		// Find
		Authenticate(ctx context.Context, dto *LoginDTO) (*User, error)

		GetUserById(ctx context.Context, id ID) (*User, error)

		SignUp(ctx context.Context, dto *SignUpDTO) (*User, error)

		GetProfile(ctx context.Context, userId ID) (*User, error)

		//UpdateUserProfile(ctx context.Context, userId ID, profile *ProfileDTO) (*User, error)
		//
		//ListAllUsers(ctx context.Context) (*[]User, error)

	}
)
