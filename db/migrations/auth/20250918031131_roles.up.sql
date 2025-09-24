BEGIN;

CREATE TABLE IF NOT EXISTS auth.roles (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(50),
    updated_at TIMESTAMP,
    updated_by VARCHAR(50),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50),

    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
    );

COMMIT;