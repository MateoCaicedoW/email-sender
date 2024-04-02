-- 20240402094056 - adding_subs_table migration
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE subs (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email text NOT NULL,
    first_name text NOT NULL,   
    last_name text NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
)