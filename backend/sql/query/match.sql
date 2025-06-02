-- name: GetMatch :one
SELECT * FROM matches
WHERE id = $1;

-- name: GetMatchesByMatchday :many
SELECT 
    m.id AS match_id,
    m.home_team_id AS home_team_id,
    m.away_team_id AS away_team_id,
    m.matchday AS matchday,
    m.match_date AS match_date,
    m.home_goals AS home_goals,
    m.away_goals AS away_goals,
    m.has_finished AS has_finished,
    hteam.long_name AS home_team_long_name,
    hteam.short_name AS home_team_short_name,
    hteam.tla AS home_team_tla,
    hteam.crest_url AS home_team_crest_url,
    ateam.long_name AS away_team_long_name,
    ateam.short_name AS away_team_short_name,
    ateam.tla AS away_team_tla,
    ateam.crest_url AS away_team_crest_url
FROM matches AS m
JOIN teams AS hteam
ON hteam.id = m.home_team_id
JOIN teams AS ateam
ON ateam.id = m.away_team_id
WHERE matchday = $1;

-- name: CreateMatch :one
INSERT INTO matches (
    home_team_id, away_team_id, matchday, match_date
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateMatch :one
UPDATE matches 
SET
    matchday = $2,
    match_date = $3,
    home_goals = $4,
    away_goals = $5,
    has_finished = "TRUE"
WHERE id = $1
RETURNING *;