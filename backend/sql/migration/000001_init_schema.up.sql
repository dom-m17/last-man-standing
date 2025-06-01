BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "comp_status" AS ENUM (
  'open',
  'in_progress',
  'complete'
);

CREATE TYPE "entry_status" AS ENUM (
  'active',
  'eliminated',
  'winner'
);

CREATE TABLE "teams" (
  "id" bigint PRIMARY KEY,
  "long_name" text UNIQUE NOT NULL,
  "short_name" text UNIQUE NOT NULL,
  "tla" text UNIQUE NOT NULL,
  "crest_url" text
);

CREATE TABLE "users" (
  "id" text DEFAULT concat('user_', uuid_generate_v4()) PRIMARY KEY,
  "username" text UNIQUE NOT NULL,
  "hashed_password" text NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "email" text UNIQUE NOT NULL,
  "phone_number" text UNIQUE,
  "date_of_birth" DATE NOT NULL,
  "favourite_team" bigint REFERENCES teams,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "competitions" (
  "id" text DEFAULT concat('comp_', uuid_generate_v4()) PRIMARY KEY,
  "name" text NOT NULL,
  "start_matchday" int NOT NULL,
  "status" comp_status DEFAULT 'open',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- TODO: Change home_team and away_team to home_team_id and away_team_id
CREATE TABLE "matches" (
  "id" bigint PRIMARY KEY,
  "home_team" bigint NOT NULL REFERENCES teams,
  "away_team" bigint NOT NULL REFERENCES teams,
  "matchday" int NOT NULL,
  "match_date" timestamptz NOT NULL,
  "home_goals" int,
  "away_goals" int,
  "has_finished" bool NOT NULL DEFAULT false
);

CREATE TABLE "entries" (
  "id" text DEFAULT concat('entry_', uuid_generate_v4()) PRIMARY KEY,
  "user_id" text NOT NULL REFERENCES users,
  "competition_id" text NOT NULL REFERENCES competitions,
  "status" entry_status NOT NULL DEFAULT 'active',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "selections" (
  "id" text DEFAULT concat('selection_', uuid_generate_v4()) PRIMARY KEY,
  "entry_id" text NOT NULL REFERENCES entries,
  "match_id" bigint NOT NULL REFERENCES matches,
  "team_id" bigint NOT NULL REFERENCES teams,
  "is_correct" bool,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "competition_matches" (
  "competition_id" text NOT NULL REFERENCES competitions,
  "match_id" bigint NOT NULL REFERENCES matches,
  PRIMARY KEY ("competition_id", "match_id")
);

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "entries" ("competition_id");

CREATE UNIQUE INDEX ON "entries" ("user_id", "competition_id");

CREATE INDEX ON "selections" ("entry_id");

CREATE UNIQUE INDEX ON "selections" ("entry_id", "team_id");

COMMENT ON COLUMN "matches"."matchday" IS 'between 1 and 38';

COMMIT;