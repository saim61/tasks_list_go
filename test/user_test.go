package test

import (
	"testing"

	"github.com/saim61/tasks_list_go/db"
	"github.com/saim61/tasks_list_go/user"
	"github.com/stretchr/testify/assert"
)

// Passing scenario for register user
func TestRegisterUser(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	userArg := user.UserRequest{
		Email:    "test@test.com",
		Password: "test@test.com",
	}

	_, _, status := user.RegisterUser(userArg, database)

	assert.Equal(t, status, true)
}

// Passing scenario for edit user
func TestEditUser(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	_, _, userDB, status := user.GetUser("test@test.com", database)

	// User found, now we can edit that user.
	assert.Equal(t, status, true)
	assert.NotEmpty(t, userDB)

	userArg := user.UserRequest{
		Email:    "test1@test.com",
		Password: "test1@test.com",
	}

	_, _, status = user.EditUser(userArg, userDB.Email, database)

	// User successfully edited
	assert.Equal(t, status, true)
}

// Fail scenario for register user
func TestFailRegisterUser(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	// Invalid arguments
	userArg := user.UserRequest{
		Email:    "test",
		Password: "",
	}

	_, _, status := user.RegisterUser(userArg, database)

	assert.Equal(t, status, false)
}

// Fail scenario for edit user
func TestFailEditUser(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	_, _, userDB, status := user.GetUser("test@test.com", database)

	// User found, now we can edit that user.
	assert.Equal(t, status, true)
	assert.NotEmpty(t, userDB)

	// Invalid edit user params
	userArg := user.UserRequest{
		Email:    "test",
		Password: "",
	}

	_, _, status = user.EditUser(userArg, userDB.Email, database)

	// User failed to edited
	assert.Equal(t, status, false)
}
