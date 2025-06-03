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
	CompStatusOpen       CompStatus = "OPEN"
	CompStatusInProgress CompStatus = "IN_PROGRESS"
	CompStatusComplete   CompStatus = "COMPLETE"
)

type Competition struct {
	ID            string `json:"id"`
	Name          string
	StartMatchday int
	Status        CompStatus
}

// User models
type User struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	DateOfBirth   time.Time `json:"dateOfBirth"`
	FavouriteTeam *string   `json:"favouriteTeam,omitempty"`
}

// Match models
type Matches []*Match

type Match struct {
	ID          string
	HomeTeam    Team
	AwayTeam    Team
	Matchday    int
	MatchDate   time.Time
	HomeGoals   int
	AwayGoals   int
	HasFinished bool
}
