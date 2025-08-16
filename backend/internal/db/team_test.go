package db

import (
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/peterldowns/testy/check"
	"github.com/stretchr/testify/require"
)

func Test_CreateTeam(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    CreateTeamParams
		wantErr func(t *testing.T, err error)
	}{
		{
			name: "success",
			args: CreateTeamParams{
				ID:        gofakeit.UUID(),
				LongName:  gofakeit.LetterN(10),
				ShortName: gofakeit.LetterN(5),
				Tla:       gofakeit.LetterN(3),
				CrestUrl:  sql.NullString{String: gofakeit.URL(), Valid: true},
			},
			wantErr: func(t *testing.T, err error) {
				t.Helper()

				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			q, close := NewTestQuerier(t)
			defer close()

			team, err := q.CreateTeam(ctx, tt.args)
			tt.wantErr(t, err)
			check.Equal(t, team.ID, tt.args.ID)
			check.Equal(t, team.LongName, tt.args.LongName)
			check.Equal(t, team.ShortName, tt.args.ShortName)
			check.Equal(t, team.Tla, tt.args.Tla)
			check.Equal(t, team.CrestUrl, tt.args.CrestUrl)
		})
	}
}
