package service

import "errors"

var ErrForeignKey = errors.New("Error invalid foreign key")
var ErrParseUUID = errors.New("Error parsing uuid value")
