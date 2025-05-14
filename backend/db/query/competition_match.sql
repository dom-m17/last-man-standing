-- name: CreateCompetitionMatch :exec
INSERT INTO competition_matches (
    competition_id, match_id
) VALUES (
    $1, $2
);