-- 20240408090229 - adding_users_management migration
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE companies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name text NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE user_companies (
    user_id UUID NOT NULL,
    company_id UUID NOT NULL,
    PRIMARY KEY (user_id, company_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

DELETE FROM subs;
DELETE FROM emails;

ALTER TABLE subs ADD COLUMN company_id UUID;
ALTER TABLE subs ADD FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;

ALTER TABLE emails ADD COLUMN company_id UUID;
ALTER TABLE emails ADD FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE;