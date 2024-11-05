-- migrate:up
ALTER TABLE users
ADD COLUMN public_id UUID DEFAULT gen_random_uuid() NOT NULL;


-- migrate:down

ALTER TABLE users
DROP COLUMN public_id;
