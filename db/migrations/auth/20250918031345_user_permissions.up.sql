BEGIN;

CREATE TABLE IF NOT EXISTS auth.user_permissions (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(50),
    updated_at TIMESTAMP,
    updated_by VARCHAR(50),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50),

    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    permission_id uuid REFERENCES permissions(id) ON DELETE CASCADE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('ALLOW','DENY'))
    );

COMMIT;