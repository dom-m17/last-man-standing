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

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar UNIQUE,
  "favourite_team" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "teams" (
  "id" smallserial PRIMARY KEY,
  "long_name" varchar UNIQUE NOT NULL,
  "short_name" varchar UNIQUE NOT NULL,
  "tla" varchar UNIQUE NOT NULL,
  "crest_url" text
);

CREATE TABLE "competitions" (
  "id" SMALLSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "start_date" date NOT NULL,
  "start_matchday" int NOT NULL,
  "status" comp_status,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "matches" (
  "id" serial PRIMARY KEY,
  "home_team" bigint NOT NULL,
  "away_team" bigint NOT NULL,
  "matchday" int NOT NULL,
  "match_date" timestamp NOT NULL,
  "home_goals" int,
  "away_goals" int,
  "has_finished" bool NOT NULL DEFAULT false
);

CREATE TABLE "entries" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "competition_id" bigint NOT NULL,
  "status" entry_status NOT NULL DEFAULT 'active'
);

CREATE TABLE "selections" (
  "id" SERIAL PRIMARY KEY,
  "entry_id" bigint NOT NULL,
  "match_id" bigint NOT NULL,
  "team_id" bigint NOT NULL,
  "is_correct" bool,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "competition_matches" (
  "competition_id" bigint NOT NULL,
  "match_id" bigint NOT NULL,
  PRIMARY KEY (competition_id,match_id)
);

CREATE INDEX ON "entries" ("user_id");

CREATE INDEX ON "entries" ("competition_id");

CREATE UNIQUE INDEX ON "entries" ("user_id", "competition_id");

CREATE INDEX ON "selections" ("entry_id");

CREATE UNIQUE INDEX ON "selections" ("entry_id", "team_id");

COMMENT ON COLUMN "matches"."matchday" IS 'between 1 and 38';

ALTER TABLE "users" ADD FOREIGN KEY ("favourite_team") REFERENCES "teams" ("id");

ALTER TABLE "matches" ADD FOREIGN KEY ("home_team") REFERENCES "teams" ("id");

ALTER TABLE "matches" ADD FOREIGN KEY ("away_team") REFERENCES "teams" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("competition_id") REFERENCES "competitions" ("id");

ALTER TABLE "selections" ADD FOREIGN KEY ("entry_id") REFERENCES "entries" ("id");

ALTER TABLE "selections" ADD FOREIGN KEY ("match_id") REFERENCES "matches" ("id");

ALTER TABLE "selections" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id");

ALTER TABLE "competition_matches" ADD FOREIGN KEY ("competition_id") REFERENCES "competitions" ("id");

ALTER TABLE "competition_matches" ADD FOREIGN KEY ("match_id") REFERENCES "matches" ("id");