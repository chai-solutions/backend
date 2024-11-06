-- migrate:up
ALTER TABLE flights
ADD COLUMN status TEXT NOT NULL DEFAULT 'scheduled';

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    event_type TEXT NOT NULL,
    flight_number TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- migrate:down
ALTER TABLE flights
DROP COLUMN status;

DROP TABLE notifications;
