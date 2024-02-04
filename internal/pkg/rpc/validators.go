package rpc

import "github.com/google/uuid"

func validateUUID(str string, dest *uuid.UUID) (err error) {
	*dest, err = uuid.Parse(str)
	return err
}
