BEGIN;

-- This data should now always come from the API
-- Insert Teams
-- INSERT INTO teams (id, long_name, short_name, tla, crest_url) VALUES
-- ('1', 'Manchester United Football Club', 'Man United', 'MUN', 'https://example.com/crest/mun.png'),
-- ('2', 'Liverpool Football Club', 'Liverpool', 'LIV', 'https://example.com/crest/liv.png'),
-- ('3', 'Arsenal Football Club', 'Arsenal', 'ARS', 'https://example.com/crest/ars.png'),
-- ('4', 'Chelsea Football Club', 'Chelsea', 'CHE', 'https://example.com/crest/che.png'),
-- ('5', 'Tottenham Hotspur Football Club', 'Tottenham', 'TOT', 'https://example.com/crest/tot.png'),
-- ('6', 'Manchester City Football Club', 'Man City', 'MCI', 'https://example.com/crest/mci.png'),
-- ('7', 'Newcastle United Football Club', 'Newcastle', 'NEW', 'https://example.com/crest/new.png'),
-- ('8', 'West Ham United Football Club', 'West Ham', 'WHU', 'https://example.com/crest/whu.png'),
-- ('9', 'Leicester City Football Club', 'Leicester', 'LEI', 'https://example.com/crest/lei.png'),
-- ('10', 'Everton Football Club', 'Everton', 'EVE', 'https://example.com/crest/eve.png'),
-- ('11', 'Brighton & Hove Albion Football Club', 'Brighton', 'BHA', 'https://example.com/crest/bha.png'),
-- ('12', 'Aston Villa Football Club', 'Aston Villa', 'AVL', 'https://example.com/crest/avl.png'),
-- ('13', 'Wolverhampton Wanderers Football Club', 'Wolves', 'WOL', 'https://example.com/crest/wol.png'),
-- ('14', 'Crystal Palace Football Club', 'Crystal Palace', 'CRY', 'https://example.com/crest/cry.png');

-- Insert Users
INSERT INTO users (username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team_id)
VALUES
('jdoe', 'hashed_pw_123', 'John', 'Doe', 'jdoe@example.com', '07123456789', '1990-01-01', 61),
('asmith', 'hashed_pw_456', 'Alice', 'Smith', 'asmith@example.com', '07234567890', '1992-02-02', 62);

-- Insert Competition
INSERT INTO competitions (id, name, start_matchday, status)
VALUES
('comp_001', 'Premier League Predictor', 1, 'OPEN');

-- This data should now always come from the API
-- Insert Matches
-- INSERT INTO matches (id, home_team_id, away_team_id, matchday, match_date, home_goals, away_goals, has_finished)
-- VALUES
-- ('1001', '1', '2', '1', '2025-08-10 15:00:00+00', 2, 1, true),
-- ('1002', '3', '4', '1', '2025-08-11 15:00:00+00', 0, 0, true),
-- ('1003', '2', '3', '2', '2025-08-17 15:00:00+00', NULL, NULL, false),
-- ('1004', '4', '1', '2', '2025-08-18 15:00:00+00', NULL, NULL, false);

-- Link Matches to Competition
INSERT INTO competition_matches (competition_id, match_id) VALUES
('comp_001', '1001'),
('comp_001', '1002'),
('comp_001', '1003'),
('comp_001', '1004');

-- Insert Entries
INSERT INTO entries (id, user_id, competition_id, status)
SELECT 'entry_001', id, 'comp_001', 'ACTIVE' FROM users WHERE username = 'jdoe';
INSERT INTO entries (id, user_id, competition_id, status)
SELECT 'entry_002', id, 'comp_001', 'ACTIVE' FROM users WHERE username = 'asmith';

-- Insert Rounds
INSERT INTO rounds (id, round_number, competition_id, matchday)
VALUES 
('round_001', 1, 'comp_001', 1),
('round_002', 2, 'comp_001', 2),
('round_003', 3, 'comp_001', 3);

-- Broken, probably
-- Insert Selections
-- INSERT INTO selections (id, entry_id, round_id, match_id, team_id, is_correct)
-- VALUES
-- ('selection_001', 'entry_001', 'round_001', '1001', '1', true), -- John chose Man United who won
-- ('selection_002', 'entry_001', 'round_002', '1002', '3', false), -- John chose Arsenal, game drew
-- ('selection_003', 'entry_002', 'round_001', '1001', '2', false), -- Alice chose Liverpool who lost
-- ('selection_004', 'entry_002', 'round_002', '1002', '4', false); -- Alice chose Chelsea, game drew

COMMIT;