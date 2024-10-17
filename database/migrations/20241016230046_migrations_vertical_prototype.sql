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

CREATE TABLE IF NOT EXISTS flight_plans (
  id INT PRIMARY KEY,
  users INT NOT NULL
);

CREATE TABLE IF NOT EXISTS flight_plan_flights (
  id INT PRIMARY KEY,
  flight_plan INT NOT NULL,
  flight INT NOT NULL
);

CREATE TABLE IF NOT EXISTS gates (
  id INT PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL,
  terminal INT NOT NULL
);

CREATE TABLE IF NOT EXISTS airports (
  id SERIAL PRIMARY KEY,
  iata VARCHAR(3) NOT NULL,
  name VARCHAR(50) NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS terminals (
  id SERIAL PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  airport INT NOT NULL,
  latitude FLOAT8 NOT NULL,
  longitude FLOAT8 NOT NULL
);

CREATE TABLE IF NOT EXISTS flights (
  id SERIAL PRIMARY KEY,
  flight_number VARCHAR NOT NULL,
  airline INT NOT NULL,
  dep_airport INT NOT NULL,
  arr_airport INT NOT NULL,
  sched_dep_time TIMESTAMP NOT NULL,
  sched_arr_time TIMESTAMP NOT NULL,
  actual_dep_time TIMESTAMP NOT NULL,
  actual_arr_time TIMESTAMP NOT NULL,
  delay_status delay_status_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS airlines (
  id SERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  iata VARCHAR(5) NOT NULL
);

ALTER TABLE flight_plans ADD FOREIGN KEY (users) REFERENCES users (id);

ALTER TABLE flight_plan_flights ADD FOREIGN KEY (flight_plan) REFERENCES flight_plans (id);

ALTER TABLE flight_plan_flights ADD FOREIGN KEY (flight) REFERENCES flights (id);

ALTER TABLE gates ADD FOREIGN KEY (terminal) REFERENCES terminals (id);

ALTER TABLE terminals ADD FOREIGN KEY (airport) REFERENCES airports (id);

ALTER TABLE flights ADD FOREIGN KEY (airline) REFERENCES airlines (id);

ALTER TABLE flights ADD FOREIGN KEY (dep_airport) REFERENCES airports (id);

ALTER TABLE flights ADD FOREIGN KEY (arr_airport) REFERENCES airports (id);

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
