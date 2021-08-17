package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var momentFixture = struct {
	MomentID    string
	UserID      string
	Name        string
	Description string
	Tags        []string
}{"123", "234", "name", "description", []string{"brutal", "hash", "tag"}}

func createMomentFromFixture() *Moment {
	return NewMoment(
		*NewMomentID(momentFixture.MomentID),
		*NewUserID(momentFixture.UserID),
		momentFixture.Name,
		momentFixture.Description,
		momentFixture.Tags,
	)
}

func TestNewMomentID(t *testing.T) {
	assert.IsType(t, &MomentID{}, NewMomentID("123"))
}

func TestMomentID_Equals(t *testing.T) {
	assert.Equal(t, true, NewMomentID("123").Equals(NewMomentID("123")))
	assert.Equal(t, false, NewMomentID("123").Equals(NewMomentID("124")))
}

func TestMomentID_Hash(t *testing.T) {
	assert.Equal(t, "123", NewMomentID("123").Hash())
}

func TestNewMoment(t *testing.T) {
	assert.IsType(t, &Moment{}, createMomentFromFixture())
}

func TestMoment_ID(t *testing.T) {
	assert.Equal(t, momentFixture.MomentID, createMomentFromFixture().MomentID().Hash())
}

func TestMoment_UserID(t *testing.T) {
	assert.Equal(t, momentFixture.UserID, createMomentFromFixture().UserID().Hash())
}

func TestMoment_OccurredAt(t *testing.T) {
	assert.IsType(t, time.Time{}, createMomentFromFixture().OccurredAt())
}

func TestMoment_Name(t *testing.T) {
	assert.Equal(t, momentFixture.Name, createMomentFromFixture().Name())
}

func TestMoment_Description(t *testing.T) {
	assert.Equal(t, momentFixture.Description, createMomentFromFixture().Description())
}

func TestMoment_Tags(t *testing.T) {
	assert.Equal(t, momentFixture.Tags, createMomentFromFixture().Tags())
}

func TestMomentCollection_Chunks(t *testing.T) {
	momentCollection := MomentCollection{*createMomentFromFixture(), *createMomentFromFixture(), *createMomentFromFixture()}
	assert.Equal(t, 2, len(momentCollection.Chunks(2)))
	assert.Equal(t, 3, len(momentCollection.Chunks(1)))
	assert.Equal(t, 1, len(momentCollection.Chunks(3)))
	assert.Equal(t, 1, len(momentCollection.Chunks(4)))
}

func TestMomentCollection_ToMap(t *testing.T) {
	momentCollection := MomentCollection{
		*createMomentFromFixture(),
		*createMomentFromFixture(),
		*NewMoment(*NewMomentID("124"), *NewUserID("234"), "name", "description", []string{}),
	}
	assert.Equal(t, 2, len(momentCollection.ToMap()))
}
