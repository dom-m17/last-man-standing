BEGIN;

DROP TABLE IF EXISTS selections;

DROP TABLE IF EXISTS entries;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS competition_matches;

DROP TABLE IF EXISTS rounds;

DROP TABLE IF EXISTS competitions;

DROP TABLE IF EXISTS matches;

DROP TABLE IF EXISTS teams;

DROP TABLE IF EXISTS refresh_tokens;

DROP TYPE IF EXISTS "comp_status";

DROP TYPE IF EXISTS "entry_status";

DROP TYPE IF EXISTS "match_status";

DROP TYPE IF EXISTS "round_status";

DROP EXTENSION IF EXISTS "uuid-ossp";

DROP FUNCTION IF EXISTS updated_at_now () CASCADE;

COMMIT;