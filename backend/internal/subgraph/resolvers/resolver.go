package graphresolvers

import (
	"github.com/dom-m17/lms/backend/internal/competition"
	"github.com/dom-m17/lms/backend/internal/entry"
	"github.com/dom-m17/lms/backend/internal/match"
	"github.com/dom-m17/lms/backend/internal/selection"
	"github.com/dom-m17/lms/backend/internal/team"
	"github.com/dom-m17/lms/backend/internal/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService        user.ServiceInterface
	CompetitionService competition.ServiceInterface
	MatchService       match.ServiceInterface
	SelectionService   selection.ServiceInterface
	TeamService        team.ServiceInterface
	EntryService       entry.ServiceInterface
}
