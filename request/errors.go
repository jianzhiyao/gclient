package request

import "errors"

var (
	ErrCanNotMarshal = errors.New(`can't marshal (implement encoding.BinaryMarshaler)`)
)
