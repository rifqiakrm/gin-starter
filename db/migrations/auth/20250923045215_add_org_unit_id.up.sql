BEGIN;

ALTER TABLE auth.users
    ADD COLUMN IF NOT EXISTS org_unit_id uuid,
    ADD CONSTRAINT fk_users_org_unit
    FOREIGN KEY (org_unit_id)
    REFERENCES org_units(id);

COMMIT;