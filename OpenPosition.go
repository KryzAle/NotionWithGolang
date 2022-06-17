package main

import (
	"errors"
)

type OpenPosition struct {
	ID    uint
	Name  string
	Skill []Skill
}

func (openPosition *OpenPosition) Init(name string, skill []Skill) (*OpenPosition, error) {

	if name == "" {
		return openPosition, errors.New("name can't be empty string")
	}
	openPosition.Name = name
	openPosition.Skill = skill

	return openPosition, nil
}
