package main

import "time"

type Participant struct {
	Id        			int       `json:"id"`
	Name      			string    `json:"name"`
	ActivationStatus 	bool      `json:"activationStatus"`
	DateOfBirth       	time.Time `json:"dateOfBirth"`
	PhoneNumber       	string 	  `json:"phoneNumber"`
	DependentParticipants []DependentParticipant  `json:"dependentParticiants"`
}

type Participants []Participant
