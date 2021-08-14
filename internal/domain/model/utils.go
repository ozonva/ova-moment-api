package model

type hasher interface {
	Hash() string
}

type domainObjectID string

func newDomainObjectID(rawID string) domainObjectID {
	return domainObjectID(rawID)
}

func (doID domainObjectID) Hash() string {
	return string(doID)
}

func (doID domainObjectID) Equals(_doID hasher) bool {
	return doID.Hash() == _doID.Hash()
}
