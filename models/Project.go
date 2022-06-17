package models

import (
	"errors"
	"time"
)

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Member      []Member `json:"member"`
	ReportingTo string   `json:"reporting_to"`
	Status      string   `json:"status"`
}

func (project *Project) Init(name string, member []Member, startDate time.Time, reportingTo string, status string) (*Project, error) {

	if name == "" {
		return nil, errors.New("name can't be empty string")
	}

	if startDate.IsZero() {
		return nil, errors.New("startDate can't be zero")
	}

	if reportingTo == "" {
		return nil, errors.New("reportingTo can't be empty string")
	}

	if status == "" {
		return nil, errors.New("status can't be empty string")
	}

	project.Name = name
	project.Member = member
	project.ReportingTo = reportingTo
	project.Status = status

	return project, nil
}
