package graphconverters

import (
	"github.com/dom-m17/lms/backend/internal/models"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

// Team Converters
func ConvertModelTeamsToGraphTeams(teams models.Teams) []*graphmodels.Team {
	gTeams := make([]*graphmodels.Team, len(teams))

	for i := range teams {
		gTeams[i] = ConvertModelTeamToGraphTeam(*teams[i])
	}

	return gTeams
}

func ConvertModelTeamToGraphTeam(team models.Team) *graphmodels.Team {
	return &graphmodels.Team{
		ID:        team.ID,
		LongName:  team.LongName,
		ShortName: team.ShortName,
		Tla:       team.Tla,
		CrestURL:  team.CrestURL,
	}
}

// Competition Converters
func ConvertModelCompetitionToGraphCompetition(competition models.Competition) *graphmodels.Competition {
	return &graphmodels.Competition{
		ID:            competition.ID,
		Name:          competition.Name,
		StartMatchday: int32(competition.StartMatchday),
		Status:        graphmodels.CompStatus(competition.Status),
	}
}

// User Converters
func ConvertModelUsersToGraphUsers(users models.Users) []*graphmodels.User {
	gUsers := make([]*graphmodels.User, len(users))
	for i := range users {
		gUsers[i] = ConvertModelUserToGraphUser(*users[i])
	}

	return gUsers
}

func ConvertModelUserToGraphUser(user models.User) *graphmodels.User {
	var favouriteTeam *graphmodels.Team
	if user.FavouriteTeam != nil {
		favouriteTeam = &graphmodels.Team{ID: *user.FavouriteTeam}
	}

	return &graphmodels.User{
		ID:            user.ID,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		PhoneNumber:   &user.PhoneNumber,
		DateOfBirth:   user.DateOfBirth.String(),
		FavouriteTeam: favouriteTeam,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

// Match Converters
func ConvertModelMatchesToGraphMatches(matches models.Matches) []*graphmodels.Match {
	gMatches := make([]*graphmodels.Match, len(matches))

	for i := range matches {
		gMatches[i] = ConvertModelMatchToGraphMatch(*matches[i])
	}

	return gMatches
}

func ConvertModelMatchToGraphMatch(match models.Match) *graphmodels.Match {
	homeGoals := int32(match.HomeGoals)
	awayGoals := int32(match.AwayGoals)

	return &graphmodels.Match{
		ID: match.ID,
		HomeTeam: &graphmodels.Team{
			ID: match.HomeTeam.ID,
		},
		AwayTeam: &graphmodels.Team{
			ID: match.AwayTeam.ID,
		},
		Matchday:    int32(match.Matchday),
		MatchDate:   match.MatchDate,
		HomeGoals:   &homeGoals,
		AwayGoals:   &awayGoals,
		HasFinished: match.HasFinished,
	}
}

// Entry Converters

// Selection Converters

// Round Converters
