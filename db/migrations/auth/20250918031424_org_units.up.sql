BEGIN;

CREATE TABLE IF NOT EXISTS auth.org_units (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT,

    name VARCHAR(100) NOT NULL,
    parent_id uuid REFERENCES org_units(id)
    );

COMMIT;