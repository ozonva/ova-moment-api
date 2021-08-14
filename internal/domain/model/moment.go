package model

import "time"

type MomentID struct {
	domainObjectID
}

func NewMomentID(ID string) *MomentID {
	return &MomentID{domainObjectID: newDomainObjectID(ID)}
}

type MomentRepository interface {
	NextID() *MomentID
	save(m *Moment)
}

type Moment struct {
	_ID         MomentID
	userID      UserID
	occurredAt  time.Time
	name        string
	description string
	tags        []string
}

func NewMoment(ID MomentID, userID UserID, name string, description string) *Moment {
	return &Moment{_ID: ID, userID: userID, occurredAt: time.Now(), name: name, description: description}
}

func (m Moment) ID() MomentID {
	return m._ID
}

func (m Moment) UserID() UserID {
	return m.userID
}

func (m Moment) OccurredAt() time.Time {
	return m.occurredAt
}

func (m Moment) Name() string {
	return m.name
}

func (m Moment) Description() string {
	return m.description
}

func (m Moment) Tags() []string {
	return m.tags
}
