package main

import (
	"errors"
	"time"
)

type Skill struct {

	ID        	int           
	Name      	string         
	CategoryID 	uint 		 
	CreatedAt 	time.Time     
	UpdatedAt 	time.Time      
	DeletedAt 	time.Time 

}

func (skill *Skill) Init(name string, CategoryID uint, UserID int) (*Skill, error) {

	if name == "" {
		return nil, errors.New("name cannot be an empty string")
	}
	if CategoryID != uint(CategoryID) {

		return nil, errors.New("name cannot be an empty string")
	}

	skill.Name = name
	skill.CategoryID = CategoryID

	return skill, nil
}

