BEGIN;

CREATE TABLE IF NOT EXISTS auth.users (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT,

    name VARCHAR(100) NOT NULL,
    title VARCHAR(50),
    email VARCHAR(150) UNIQUE NOT NULL,
    phone_number VARCHAR(13),
    dob DATE,
    photo TEXT,
    password TEXT NOT NULL,
    forgot_password_token TEXT,
    otp VARCHAR(6),
    status VARCHAR(50) NOT NULL
);

COMMIT;