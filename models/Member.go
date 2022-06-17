package models

import (
	"errors"
)

type Member struct {
	ID           uint
	NameUser     string
	NamePosition string
}

func (member *Member) Init(nameUser, namePosition string) (*Member, error) {

	if nameUser == "" {
		return nil, errors.New("name_user cannot be an empty string")
	}

	if namePosition == "" {
		return nil, errors.New("name_position cannot be an empty string")
	}

	member.NameUser = nameUser
	member.NamePosition = namePosition
	return member, nil
}
