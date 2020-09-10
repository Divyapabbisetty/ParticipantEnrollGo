package main

import "fmt"

var currentId int

var participants Participants

// Give us some seed data
func init() {
	RepoCreateParticipant(Participant{Name: "Write presentation"})
	RepoCreateParticipant(Participant{Name: "Host meetup"})
}

func RepoFindParticipant(id int) Participant {
	for _, t := range participants {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Participant{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateParticipant(t Participant) Participant {
	currentId += 1
	t.Id = currentId
	participants = append(participants, t)
	return t
}

func RepoDestroyParticipant(id int) error {
	for i, t := range participants {
		if t.Id == id {
			participants = append(participants[:i], participants[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Participant with id of %d to delete", id)
}
