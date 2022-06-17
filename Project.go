package main

import (
	"errors"
	"time"
)

type Project struct {
	ID            string
	Name          string
	ClientId      string
	Member        []Member
	StartDate     time.Time
	OpenPositions []OpenPosition
	Description   string
	Status        string
}

func (project *Project) Init(name string, clientId string, member []Member, startDate time.Time, openposition []OpenPosition, description string, status string) (*Project, error) {

	if name == "" {
		return nil, errors.New("name can't be empty string")
	}
	if clientId == "" {
		return nil, errors.New("clientId can't be empty string")
	}

	if startDate.IsZero() {
		return nil, errors.New("startDate can't be zero")
	}

	if description == "" {
		return nil, errors.New("description can't be empty string")
	}

	if status == "" {
		return nil, errors.New("status can't be empty string")
	}

	project.Name = name
	project.ClientId = clientId
	project.Member = member
	project.OpenPositions = openposition
	project.StartDate = startDate
	project.Description = description
	project.Status = status

	return project, nil
}
