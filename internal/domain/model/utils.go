package model

type hasher interface {
	Hash() string
}

type domainObjectID string

func newDomainObjectID(rawID string) domainObjectID {
	return domainObjectID(rawID)
}

func (ID domainObjectID) Hash() string {
	return string(ID)
}

func (ID domainObjectID) Equals(_ID hasher) bool {
	return ID.Hash() == _ID.Hash()
}
