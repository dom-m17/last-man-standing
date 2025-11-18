BEGIN;

-- Clear data from tables
DELETE FROM selections;

DELETE FROM rounds;

DELETE FROM entries;

DELETE FROM competition_matches;

DELETE FROM competitions;

DELETE FROM users;

-- Insert Users
INSERT INTO
    users (
        username,
        hashed_password,
        first_name,
        last_name,
        email,
        phone_number,
        date_of_birth,
        favourite_team_id
    )
VALUES
    (
        'a.altman',
        'hashed_pw_123',
        'Alan',
        'Altman',
        'a.altman@example.com',
        '07123456789',
        '1990-01-01',
        61
    ),
    (
        'b.brown',
        'hashed_pw_456',
        'Bob',
        'Brown',
        'b.brown@example.com',
        '07234567890',
        '1992-02-02',
        62
    ),
    (
        'c.cromwell',
        'hashed_pw_789',
        'Chris',
        'Cromwell',
        'c.cromwell@example.com',
        '07987654321',
        '1980-11-29',
        57
    ),
    (
        'd.dixon',
        'hashed_pw_abc',
        'Dean',
        'Dixon',
        'd.dixon@example.com',
        '07111213145',
        '1999-06-12',
        58
    );

-- Insert Competition
INSERT INTO
    competitions (id, name, start_matchday, status)
VALUES
    ('comp_001', 'Competition 1', 1, 'COMPLETED'),
    ('comp_002', 'Competition 2', 5, 'IN_PROGRESS'),
    ('comp_003', 'Competition 3', 10, 'OPEN');

-- Link Matches to Competition
INSERT INTO
    competition_matches (competition_id, match_id)
VALUES
    ('comp_001', 537785),
    ('comp_001', 537786),
    ('comp_001', 537787),
    ('comp_001', 537789),
    ('comp_001', 537790),
    ('comp_001', 537791),
    ('comp_001', 537788),
    ('comp_001', 537792),
    ('comp_001', 537793),
    ('comp_001', 537794),
    ('comp_001', 537804),
    ('comp_001', 537802),
    ('comp_001', 537795),
    ('comp_001', 537798),
    ('comp_001', 537799),
    ('comp_001', 537797),
    ('comp_001', 537796),
    ('comp_001', 537800),
    ('comp_001', 537801),
    ('comp_001', 537803),
    ('comp_001', 537808),
    ('comp_001', 537805),
    ('comp_001', 537811),
    ('comp_001', 537813),
    ('comp_001', 537814),
    ('comp_001', 537810),
    ('comp_001', 537807),
    ('comp_001', 537812),
    ('comp_001', 537809),
    ('comp_001', 537806),
    ('comp_001', 537817),
    ('comp_001', 537815),
    ('comp_001', 537816),
    ('comp_001', 537820),
    ('comp_001', 537821),
    ('comp_001', 537823),
    ('comp_001', 537824),
    ('comp_001', 537818),
    ('comp_001', 537819),
    ('comp_001', 537822),
    ('comp_001', 537831),
    ('comp_001', 537827),
    ('comp_001', 537829),
    ('comp_001', 537833),
    ('comp_001', 537834),
    ('comp_001', 537832),
    ('comp_001', 537830),
    ('comp_001', 537825),
    ('comp_001', 537826),
    ('comp_001', 537828),
    ('comp_002', 537831),
    ('comp_002', 537827),
    ('comp_002', 537829),
    ('comp_002', 537833),
    ('comp_002', 537834),
    ('comp_002', 537832),
    ('comp_002', 537830),
    ('comp_002', 537825),
    ('comp_002', 537826),
    ('comp_002', 537828),
    ('comp_002', 537837),
    ('comp_002', 537836),
    ('comp_002', 537838),
    ('comp_002', 537840),
    ('comp_002', 537841),
    ('comp_002', 537843),
    ('comp_002', 537844),
    ('comp_002', 537835),
    ('comp_002', 537842),
    ('comp_002', 537839),
    ('comp_002', 537845),
    ('comp_002', 537851),
    ('comp_002', 537847),
    ('comp_002', 537852),
    ('comp_002', 537849),
    ('comp_002', 537846),
    ('comp_002', 537850),
    ('comp_002', 537853),
    ('comp_002', 537854),
    ('comp_002', 537848),
    ('comp_002', 537862),
    ('comp_002', 537856),
    ('comp_002', 537857),
    ('comp_002', 537858),
    ('comp_002', 537861),
    ('comp_002', 537855),
    ('comp_002', 537859),
    ('comp_002', 537863),
    ('comp_002', 537860),
    ('comp_002', 537864),
    ('comp_002', 537871),
    ('comp_002', 537869),
    ('comp_002', 537873),
    ('comp_002', 537872),
    ('comp_002', 537868),
    ('comp_002', 537865),
    ('comp_002', 537866),
    ('comp_002', 537867),
    ('comp_002', 537874),
    ('comp_002', 537870);

-- Insert Entries
INSERT INTO
    entries (id, user_id, competition_id, status)
SELECT
    'entry_001',
    id,
    'comp_001',
    'ACTIVE'
FROM
    users
WHERE
    username = 'a.altman';

INSERT INTO
    entries (id, user_id, competition_id, status)
SELECT
    'entry_002',
    id,
    'comp_001',
    'ACTIVE'
FROM
    users
WHERE
    username = 'b.brown';

-- Insert Rounds
INSERT INTO
    rounds (id, round_number, competition_id, matchday)
VALUES
    ('round_001', 1, 'comp_001', 1),
    ('round_002', 2, 'comp_001', 2),
    ('round_003', 3, 'comp_001', 3);

--! Broken
-- Insert Selections
-- INSERT INTO selections (id, entry_id, round_id, match_id, team_id, is_correct)
-- VALUES
-- ('selection_001', 'entry_001', 'round_001', '1001', '1', true), -- John chose Man United who won
-- ('selection_002', 'entry_001', 'round_002', '1002', '3', false), -- John chose Arsenal, game drew
-- ('selection_003', 'entry_002', 'round_001', '1001', '2', false), -- Alice chose Liverpool who lost
-- ('selection_004', 'entry_002', 'round_002', '1002', '4', false); -- Alice chose Chelsea, game drew
COMMIT;