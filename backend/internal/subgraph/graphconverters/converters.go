package graphconverters

import (
	"github.com/dom-m17/lms/backend/internal/models"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

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
			ID:        match.HomeTeam.ID,
			LongName:  match.HomeTeam.LongName,
			ShortName: match.HomeTeam.ShortName,
			Tla:       match.HomeTeam.Tla,
			CrestURL:  match.HomeTeam.CrestURL,
		},
		AwayTeam: &graphmodels.Team{
			ID:        match.AwayTeam.ID,
			LongName:  match.AwayTeam.LongName,
			ShortName: match.AwayTeam.ShortName,
			Tla:       match.AwayTeam.Tla,
			CrestURL:  match.AwayTeam.CrestURL,
		},
		Matchday:    int32(match.Matchday),
		MatchDate:   match.MatchDate,
		HomeGoals:   &homeGoals,
		AwayGoals:   &awayGoals,
		HasFinished: match.HasFinished,
	}
}
