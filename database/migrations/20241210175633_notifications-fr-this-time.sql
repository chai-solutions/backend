-- migrate:up
DROP TABLE IF EXISTS notifications;
CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  message TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  "user" INT NOT NULL,

  CONSTRAINT "notifications-user-fkey"
    FOREIGN KEY ("user")
      REFERENCES users (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

-- migrate:down
DROP TABLE notifications;
