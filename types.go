package main

// DateFormat ...
const DateFormat = "2006-01-02"

// ChaptersNumber ...
const ChaptersNumber = 30

// Payload will be received from the user
type Payload struct {
	StartDate      string `json:"start_date" form:"start_date" query:"start_date"`
	Days           string `json:"days" form:"days" query:"days"`
	SessionsNumber int    `json:"sessions_number" form:"sessions_number" query:"sessions_number"`
}

// Session ...
type Session struct {
	Day  string
	Date string
}

// Sessions ...
type Sessions []Session
