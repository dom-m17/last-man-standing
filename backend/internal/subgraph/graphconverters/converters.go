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
		ID:          match.ID,
		HomeTeam:    ConvertModelTeamToGraphTeam(match.HomeTeam),
		AwayTeam:    ConvertModelTeamToGraphTeam(match.AwayTeam),
		Matchday:    int32(match.Matchday),
		MatchDate:   match.MatchDate,
		HomeGoals:   &homeGoals,
		AwayGoals:   &awayGoals,
		HasFinished: match.HasFinished,
	}
}

// Entry Converters

// Selection Converters
