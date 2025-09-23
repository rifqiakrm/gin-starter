BEGIN;

CREATE TABLE IF NOT EXISTS auth.roles (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT,

    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
    );

COMMIT;