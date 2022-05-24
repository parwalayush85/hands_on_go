package blerr

type Kind int

const (
	KindUnkown Kind = iota
	KindNotFound
	KindInvalidInput
	KindInternalServerError
)
