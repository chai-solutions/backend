-- migrate:up

DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'delay_status_enum') THEN
    CREATE TYPE delay_status_enum AS ENUM ('on-time', 'delayed', 'cancelled');
  END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL,
  created_at TIMESTAMP DEFAULT NOW(),
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,

  CONSTRAINT "users-pkey" PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS sessions (
  id SERIAL,
  user_id INT NOT NULL,
  token VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,

  CONSTRAINT "sessions-pkey" PRIMARY KEY (id),

  CONSTRAINT "sessions-user_id-fkey"
    FOREIGN KEY (user_id)
      REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "flight_plan" (
  "id" int PRIMARY KEY,
  "users" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "flight_plan_flights" (
  "id" int PRIMARY KEY,
  "flight_plan" int NOT NULL,
  "flight" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "gates" (
  "id" int PRIMARY KEY,
  "name" varchar(20) NOT NULL,
  "latitude" float8 NOT NULL,
  "longitude" float8 NOT NULL,
  "terminal" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "airport" (
  "id" serial PRIMARY KEY,
  "iata" varchar(3) NOT NULL,
  "name" varchar(50) NOT NULL,
  "latitude" float8 NOT NULL,
  "longitude" float8 NOT NULL
);

CREATE TABLE IF NOT EXISTS "terminal" (
  "id" serial PRIMARY KEY,
  "name" varchar(20) NOT NULL,
  "airport" int NOT NULL,
  "latitude" float8 NOT NULL,
  "longitude" float8 NOT NULL
);

CREATE TABLE IF NOT EXISTS "flight" (
  "id" serial PRIMARY KEY,
  "flight_number" varchar NOT NULL,
  "airline" int NOT NULL,
  "dep_airport" int NOT NULL,
  "arr_airport" int NOT NULL,
  "sched_dep_time" timestamp NOT NULL,
  "sched_arr_time" timestamp NOT NULL,
  "actual_dep_time" timestamp NOT NULL,
  "actual_arr_time" timestamp NOT NULL,
  "delay_status" delay_status_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS "airline" (
  "id" serial PRIMARY KEY,
  "name" varchar(30) NOT NULL,
  "iata" varchar(5) NOT NULL
);

ALTER TABLE "flight_plan" ADD FOREIGN KEY ("users") REFERENCES "users" ("id");

ALTER TABLE "flight_plan_flights" ADD FOREIGN KEY ("flight_plan") REFERENCES "flight_plan" ("id");

ALTER TABLE "flight_plan_flights" ADD FOREIGN KEY ("flight") REFERENCES "flight" ("id");

ALTER TABLE "gates" ADD FOREIGN KEY ("terminal") REFERENCES "terminal" ("id");

ALTER TABLE "terminal" ADD FOREIGN KEY ("airport") REFERENCES "airport" ("id");

ALTER TABLE "flight" ADD FOREIGN KEY ("airline") REFERENCES "airline" ("id");

ALTER TABLE "flight" ADD FOREIGN KEY ("dep_airport") REFERENCES "airport" ("id");

ALTER TABLE "flight" ADD FOREIGN KEY ("arr_airport") REFERENCES "airport" ("id");

-- migrate:down
DROP TABLE IF EXISTS "flight_plan" CASCADE;
DROP TABLE IF EXISTS "flight_plan_flights" CASCADE;
DROP TABLE IF EXISTS "flight" CASCADE;
DROP TABLE IF EXISTS "terminal" CASCADE;
DROP TABLE IF EXISTS "gates" CASCADE;
DROP TABLE IF EXISTS "airport" CASCADE;
DROP TABLE IF EXISTS "airline" CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP TYPE IF EXISTS delay_status_enum;
