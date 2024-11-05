-- migrate:up
ALTER TABLE users
ADD COLUMN public_id UUID DEFAULT gen_random_uuid() NOT NULL,
ADD CONSTRAINT unique_public_id UNIQUE (public_id);

-- migrate:down
ALTER TABLE users
DROP CONSTRAINT unique_public_id;

ALTER TABLE users
DROP COLUMN public_id;
