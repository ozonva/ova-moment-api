package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserID(t *testing.T) {
	assert.IsType(t, &UserID{}, NewUserID("123"))
}

func TestUserID_Equals(t *testing.T) {
	assert.Equal(t, true, NewUserID("123").Equals(NewUserID("123")))
	assert.Equal(t, false, NewUserID("123").Equals(NewUserID("124")))
}

func TestUserID_Hash(t *testing.T) {
	assert.Equal(t, "123", NewUserID("123").Hash())
}
