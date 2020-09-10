package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ParticipantIndex",
		"GET",
		"/api/participants",
		ParticipantIndex,
	},
	Route{
		"ParticipantCreate",
		"POST",
		"/api/participants",
		ParticipantCreate,
	},
	Route{
		"ParticipantShow",
		"GET",
		"/api/participants/{participantId}",
		ParticipantShow,
	},
}
