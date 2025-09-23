BEGIN;

CREATE TABLE IF NOT EXISTS auth.user_permissions (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT,

    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('ALLOW','DENY'))
    );

COMMIT;