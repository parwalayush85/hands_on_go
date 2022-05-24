package blerr

import "errors"

type KindError struct {
	kind Kind
	err  error
}

func (ke *KindError) Error() string {
	if ke.err == nil {
		return ""
	}
	return ke.err.Error()
}
func (ke *KindError) Unwrap() error {
	return ke.err
}

func SetKind(err error, kind Kind) error {
	return &KindError{kind: kind, err: err}
}
func GetKind(err error) Kind {
	var kindErr *KindError
	if errors.As(err, &kindErr) {
		return kindErr.kind
	}
	return KindUnkown
}
