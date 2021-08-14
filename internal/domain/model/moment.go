package model

import (
	"github.com/ozonva/ova-moment-api/internal/utils"
	"time"
)

type MomentID struct {
	domainObjectID
}

func NewMomentID(rawID string) *MomentID {
	return &MomentID{domainObjectID: newDomainObjectID(rawID)}
}

type MomentRepository interface {
	NextID() *MomentID
	save(m *Moment)
}

type Moment struct {
	momentID    MomentID
	userID      UserID
	occurredAt  time.Time
	name        string
	description string
	tags        []string
}

func NewMoment(momentID MomentID, userID UserID, name string, description string) *Moment {
	return &Moment{momentID: momentID, userID: userID, occurredAt: time.Now(), name: name, description: description}
}

func (m Moment) ID() MomentID {
	return m.momentID
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

type MomentCollection []Moment

func (mc MomentCollection) Chunks(chunkSize int) [][]Moment {
	if chunkSize <= 0 {
		panic("Chunk size must be a positive integer")
	}

	chunks := make([][]Moment, (len(mc)+chunkSize-1)/chunkSize)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = mc[i*chunkSize : utils.MinInt(i*chunkSize+chunkSize, len(mc))]
	}

	return chunks
}

func (mc MomentCollection) ToMap() map[MomentID]Moment {
	m := make(map[MomentID]Moment)
	for _, v := range mc {
		m[v.ID()] = v
	}

	return m
}
