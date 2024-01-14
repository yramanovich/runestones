-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS runestones (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    url text UNIQUE,
    created_time timestamptz NOT NULL DEFAULT now()
);