-- migrate:up
CREATE TABLE users (
  id SERIAL,
  created_at TIMESTAMP DEFAULT NOW(),
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,

  CONSTRAINT "users-pkey" PRIMARY KEY (id)
);

CREATE TABLE sessions (
  id SERIAL,
  user_id INT,
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
DROP TABLE accounts;

-- migrate:down
DROP TABLE sessions;
DROP TABLE users;
