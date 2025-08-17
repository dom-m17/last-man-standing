package models

import "time"

// Team models
type Teams []*Team

type Team struct {
	ID        string
	LongName  string
	ShortName string
	Tla       string
	CrestURL  string
}

// Competition models
type CompStatus string

const (
	// Open is the status for competitions that have not yet started and are open to new entries
	CompStatusOpen CompStatus = "OPEN"
	// In Progress is the status for competitions that currently underway
	CompStatusInProgress CompStatus = "IN_PROGRESS"
	// Complete is the status for competitions that have finished
	CompStatusComplete CompStatus = "COMPLETE"
)

type Competition struct {
	ID            string `json:"id"`
	Name          string
	StartMatchday int
	Status        CompStatus
}

// User models
type Users []*User
type User struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	FavouriteTeam *string   `json:"favouriteTeam,omitempty"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Match models
type Matches []*Match

type Match struct {
	ID        string
	HomeTeam  Team
	AwayTeam  Team
	Matchday  int
	MatchDate time.Time
	HomeGoals int
	AwayGoals int
	Status    MatchStatus
}

type MatchStatus string

const (
	// Scheduled are matches that are in the future without a confirmed date/time
	MatchStatusScheduled MatchStatus = "SCHEDULED"
	// Timed are matches that are in the future and have a confirmed date/time
	MatchStatusTimed MatchStatus = "TIMED"
	// In Play are matches that are currently being played
	MatchStatusInPlay MatchStatus = "IN_PLAY"
	// Finished are matches that have been completed
	MatchStatusFinished MatchStatus = "FINISHED"
)

// Entry models
type EntryStatus string

const (
	EntryStatusEliminated EntryStatus = "ELIMINATED"
	EntryStatusActive     EntryStatus = "ACTIVE"
	EntryStatusWinner     EntryStatus = "WINNER"
)
