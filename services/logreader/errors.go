package logreader

import "errors"

var (
	ErrToReadFile   = errors.New("error to read file")
	ErrEmptyLogFile = errors.New("empty log file")
)
