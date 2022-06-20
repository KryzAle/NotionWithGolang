package models

import (
	"errors"
)

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	ClientId	string   `json:"client_id"`
	Member      []Member `json:"member"`
	ReportingTo string   `json:"reporting_to"`
	Status      string   `json:"status"`
}

func (project *Project) Init(id ,name, clientId, reportingTo, status string, member []Member) (*Project, error) {

	if id == "" {
		return nil, errors.New("id is empty")
	}
	
	if name == "" {
		return nil, errors.New("name can't be empty string")
	}

	if clientId == "" {
		return nil, errors.New("clientId can't be empty string")
	}

	if reportingTo == "" {
		return nil, errors.New("reportingTo can't be empty string")
	}

	if status == "" {
		return nil, errors.New("status can't be empty string")
	}

	project.ID = id
	project.Name = name
	project.ClientId = clientId
	project.ReportingTo = reportingTo
	project.Status = status
	project.Member = member

	return project, nil
}
