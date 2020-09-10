package main

import "time"

type DependentParticipant struct {
	Id        			int       `json:"id"`
	Name      			string    `json:"name"`
	DateOfBirth       	time.Time `json:"dateOfBirth"`
}

type DependentParticipants []DependentParticipant
