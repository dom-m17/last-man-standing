BEGIN;

CREATE
OR REPLACE FUNCTION updated_at_now () RETURNS TRIGGER AS
$$
BEGIN NEW.updated_at = now ();

RETURN NEW;

END;

$$ 
LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "comp_status" AS ENUM ('OPEN', 'IN_PROGRESS', 'COMPLETED');

CREATE TYPE "entry_status" AS ENUM ('ACTIVE', 'ELIMINATED', 'WINNER');

CREATE TYPE "match_status" AS ENUM ('FINISHED', 'IN_PLAY', 'SCHEDULED', 'TIMED');

CREATE TYPE "round_status" AS ENUM ('PENDING', 'IN_PLAY', 'FINISHED');

CREATE TABLE
  IF NOT EXISTS "teams" (
    "id" text PRIMARY KEY,
    "long_name" text UNIQUE NOT NULL,
    "short_name" text UNIQUE NOT NULL,
    "tla" text UNIQUE NOT NULL,
    "crest_url" text,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "competitions" (
    "id" text DEFAULT concat ('comp_', uuid_generate_v4 ()) PRIMARY KEY,
    "name" text NOT NULL,
    "start_matchday" int NOT NULL,
    "status" comp_status NOT NULL DEFAULT 'OPEN',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "users" (
    "id" text DEFAULT concat ('user_', uuid_generate_v4 ()) PRIMARY KEY,
    "username" text UNIQUE NOT NULL,
    "hashed_password" text NOT NULL,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "email" text UNIQUE NOT NULL,
    "phone_number" text UNIQUE,
    "date_of_birth" DATE NOT NULL,
    "favourite_team_id" text REFERENCES teams ON DELETE CASCADE,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "matches" (
    "id" text PRIMARY KEY,
    "home_team_id" text NOT NULL REFERENCES teams ON DELETE CASCADE,
    "away_team_id" text NOT NULL REFERENCES teams ON DELETE CASCADE,
    "matchday" int NOT NULL,
    "match_date" timestamptz NOT NULL,
    "home_goals" int,
    "away_goals" int,
    "status" match_status NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "entries" (
    "id" text DEFAULT concat ('entry_', uuid_generate_v4 ()) PRIMARY KEY,
    "user_id" text NOT NULL REFERENCES users,
    "competition_id" text NOT NULL REFERENCES competitions ON DELETE CASCADE,
    "status" entry_status NOT NULL DEFAULT 'ACTIVE',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "rounds" (
    "id" text DEFAULT concat ('round_', uuid_generate_v4 ()) PRIMARY KEY,
    "round_number" text NOT NULL,
    "competition_id" text NOT NULL REFERENCES competitions ON DELETE CASCADE,
    "matchday" int NOT NULL,
    "status" round_status NOT NULL DEFAULT 'PENDING',
    --! This needs to be NOT NULL, keeping it this way for now for ease of development
    "entry_deadline" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "selections" (
    "id" text DEFAULT concat ('selection_', uuid_generate_v4 ()) PRIMARY KEY,
    "entry_id" text NOT NULL REFERENCES entries ON DELETE CASCADE,
    "round_id" text NOT NULL REFERENCES rounds ON DELETE CASCADE,
    "match_id" text NOT NULL REFERENCES matches ON DELETE CASCADE,
    "team_id" text NOT NULL REFERENCES teams ON DELETE CASCADE,
    "is_correct" bool,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "competition_matches" (
    "competition_id" text NOT NULL REFERENCES competitions ON DELETE CASCADE,
    "match_id" text NOT NULL REFERENCES matches,
    PRIMARY KEY ("competition_id", "match_id"),
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "refresh_tokens" (
    "id" text DEFAULT concat ('refresh_token_', uuid_generate_v4 ()) PRIMARY KEY,
    "user_id" text NOT NULL REFERENCES users ON DELETE CASCADE,
    "token_hash" TEXT NOT NULL,
    "expires_at" timestamptz NOT NULL DEFAULT (now () + interval '30 day'),
    "revoked" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW ()
  );

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "entries" ("competition_id");

CREATE UNIQUE INDEX ON "entries" ("user_id", "competition_id");

CREATE INDEX ON "selections" ("entry_id");

CREATE UNIQUE INDEX ON "selections" ("entry_id", "team_id");

COMMENT ON COLUMN "matches"."matchday" IS 'between 1 and 38';

CREATE TRIGGER teams_updated_at BEFORE
UPDATE ON teams FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER competitions_updated_at BEFORE
UPDATE ON competitions FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER users_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER matches_updated_at BEFORE
UPDATE ON matches FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER entries_updated_at BEFORE
UPDATE ON entries FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER rounds_updated_at BEFORE
UPDATE ON rounds FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER selections_updated_at BEFORE
UPDATE ON selections FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

CREATE TRIGGER competition_matches_updated_at BEFORE
UPDATE ON competition_matches FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

COMMIT;