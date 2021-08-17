package model

type UserID struct {
	domainObjectID
}

func NewUserID(rawID string) *UserID {
	return &UserID{newDomainObjectID(rawID)}
}
