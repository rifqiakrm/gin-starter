BEGIN;

CREATE TABLE IF NOT EXISTS auth.permissions (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT,

    key VARCHAR(100) UNIQUE NOT NULL,   -- e.g. "course:view"
    description TEXT
    );

COMMIT;