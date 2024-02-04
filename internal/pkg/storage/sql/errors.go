package sql

import (
	"errors"
)

var (
	ErrEntityExist    = errors.New("Entity already exist")
	ErrEntityNotExist = errors.New("Entity doesn't exist")
	ErrForeignKey     = errors.New("Error invalid foreign key")
)
