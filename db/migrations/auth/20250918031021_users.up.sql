BEGIN;

CREATE TABLE IF NOT EXISTS auth.users (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by VARCHAR(50),
    updated_at TIMESTAMP,
    updated_by VARCHAR(50),
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50),

    employee_id VARCHAR(100),
    name VARCHAR(100) NOT NULL,
    title VARCHAR(50),
    email VARCHAR(150) UNIQUE NOT NULL,
    username VARCHAR(25) UNIQUE,
    password TEXT NOT NULL,
    phone_number VARCHAR(13),
    dob DATE,
    address TEXT,
    photo TEXT,
    forgot_password_token TEXT,
    otp VARCHAR(6),
    status VARCHAR(50) NOT NULL
);

COMMIT;