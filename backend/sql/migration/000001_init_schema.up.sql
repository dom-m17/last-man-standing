BEGIN;

CREATE OR REPLACE FUNCTION updated_at_now()
RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

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
    "favourite_team_id" text REFERENCES teams,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "matches" (
    "id" text PRIMARY KEY,
    "home_team_id" text NOT NULL REFERENCES teams,
    "away_team_id" text NOT NULL REFERENCES teams,
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
    "competition_id" text NOT NULL REFERENCES competitions,
    "status" entry_status NOT NULL DEFAULT 'ACTIVE',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "rounds" (
    "id" text DEFAULT concat ('round_', uuid_generate_v4 ()) PRIMARY KEY,
    "round_number" text NOT NULL,
    "competition_id" text NOT NULL REFERENCES competitions,
    "matchday" int NOT NULL,
    "status" round_status NOT NULL DEFAULT 'PENDING',
    "entry_deadline" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "selections" (
    "id" text DEFAULT concat ('selection_', uuid_generate_v4 ()) PRIMARY KEY,
    "entry_id" text NOT NULL REFERENCES entries,
    "round_id" text NOT NULL REFERENCES rounds,
    "match_id" text NOT NULL REFERENCES matches,
    "team_id" text NOT NULL REFERENCES teams,
    "is_correct" bool,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE TABLE
  IF NOT EXISTS "competition_matches" (
    "competition_id" text NOT NULL REFERENCES competitions,
    "match_id" text NOT NULL REFERENCES matches,
    PRIMARY KEY ("competition_id", "match_id"),
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
  );

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "entries" ("competition_id");

CREATE UNIQUE INDEX ON "entries" ("user_id", "competition_id");

CREATE INDEX ON "selections" ("entry_id");

CREATE UNIQUE INDEX ON "selections" ("entry_id", "team_id");

COMMENT ON COLUMN "matches"."matchday" IS 'between 1 and 38';

CREATE TRIGGER teams BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER competitions BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER users BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER matches BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER entries BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER rounds BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER selections BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();
CREATE TRIGGER competition_matches BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE updated_at_now ();

COMMIT;