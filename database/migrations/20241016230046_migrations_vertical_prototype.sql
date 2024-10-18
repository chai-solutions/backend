-- migrate:up

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
  token TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP NOT NULL,

  CONSTRAINT "sessions-pkey" PRIMARY KEY (id),

  CONSTRAINT "sessions-user_id-fkey"
    FOREIGN KEY (user_id)
      REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS flight_plans (
  id SERIAL PRIMARY KEY,
  users INT NOT NULL
);

CREATE TABLE IF NOT EXISTS flight_plan_flights (
  id SERIAL PRIMARY KEY,
  flight_plan INT NOT NULL,
  flight INT NOT NULL
);

CREATE TABLE IF NOT EXISTS gates (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL,
  terminal INT NOT NULL
);

CREATE TABLE IF NOT EXISTS airports (
  id SERIAL PRIMARY KEY,
  iata TEXT NOT NULL,
  name TEXT NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS terminals (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  airport INT NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS flights (
  id SERIAL PRIMARY KEY,
  flight_number TEXT NOT NULL,
  airline INT NOT NULL,
  dep_airport INT NOT NULL,
  arr_airport INT NOT NULL,
  sched_dep_time TIMESTAMP NOT NULL,
  sched_arr_time TIMESTAMP NOT NULL,
  actual_dep_time TIMESTAMP NOT NULL,
  actual_arr_time TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS airlines (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  iata TEXT NOT NULL
);

ALTER TABLE flight_plans ADD FOREIGN KEY (users) REFERENCES users (id);

ALTER TABLE flight_plan_flights ADD FOREIGN KEY (flight_plan) REFERENCES flight_plans (id);
ALTER TABLE flight_plan_flights
DROP CONSTRAINT flight_plan_flights_flight_plan_fkey,
ADD CONSTRAINT flight_plan_flights_flight_plan_fkey
FOREIGN KEY (flight_plan)
REFERENCES flight_plans (id)
ON DELETE CASCADE;

ALTER TABLE flight_plan_flights ADD FOREIGN KEY (flight) REFERENCES flights (id);

ALTER TABLE gates ADD FOREIGN KEY (terminal) REFERENCES terminals (id);

ALTER TABLE terminals ADD FOREIGN KEY (airport) REFERENCES airports (id);

ALTER TABLE flights ADD FOREIGN KEY (airline) REFERENCES airlines (id);

ALTER TABLE flights ADD FOREIGN KEY (dep_airport) REFERENCES airports (id);

ALTER TABLE flights ADD FOREIGN KEY (arr_airport) REFERENCES airports (id);

-- migrate:down
DROP TABLE IF EXISTS "flight_plans" CASCADE;
DROP TABLE IF EXISTS "flight_plan_flights" CASCADE;
DROP TABLE IF EXISTS "flights" CASCADE;
DROP TABLE IF EXISTS "terminals" CASCADE;
DROP TABLE IF EXISTS "gates" CASCADE;
DROP TABLE IF EXISTS "airports" CASCADE;
DROP TABLE IF EXISTS "airlines" CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS users CASCADE;
