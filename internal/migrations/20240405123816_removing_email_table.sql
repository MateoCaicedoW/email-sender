-- 20240405123816 - removing_email_table migration

DROP TABLE emails;

CREATE TABLE emails (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    message TEXT NOT NULL,
    sent BOOLEAN DEFAULT TRUE,
    subject TEXT NOT NULL,

    --Attachment attributes
    attachment_name TEXT,
    attachment_content BYTEA NOT NULL,
    scheduled BOOLEAN DEFAULT FALSE,
    scheduled_at TIMESTAMP,

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);