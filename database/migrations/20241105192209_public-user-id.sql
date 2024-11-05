-- migrate:up
ALTER TABLE users
ADD COLUMN uuid UUID DEFAULT gen_random_uuid() NOT NULL;


-- migrate:down

ALTER TABLE users
DROP COLUMN uuid;
