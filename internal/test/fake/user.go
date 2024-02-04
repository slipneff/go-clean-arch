package fake

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/slipneff/go-clean-arch/internal/pkg/model"
)

func User() model.User {
	return model.User{
		Id:   uuid.New(),
		Name: randomdata.RandStringRunes(15),
	}
}
