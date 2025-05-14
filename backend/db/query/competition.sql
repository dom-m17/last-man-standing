-- name: GetCompetition :one
SELECT * FROM competitions
WHERE id = $1;

-- name: CreateCompetition :one
INSERT INTO competitions (
    name,
    start_matchday
) VALUES (
    $1, $2
)
RETURNING *;
