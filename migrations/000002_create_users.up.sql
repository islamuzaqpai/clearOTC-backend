CREATE TABLE clearotc.users
(
    id             UUID PRIMARY KEY,
    name           VARCHAR(255) NOT NULL,
    lastname       VARCHAR(255) NOT NULL,
    email          VARCHAR(255) NOT NULL,
    password_hash  TEXT         NOT NULL,
    wallet_address VARCHAR(42)  NULL,
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT now(),
    deleted_at     TIMESTAMPTZ  NULL
);

CREATE UNIQUE INDEX users_email_unique_idx
    ON clearotc.users (LOWER(email));

CREATE UNIQUE INDEX users_wallet_unique_idx
    ON clearotc.users (LOWER(wallet_address))
    WHERE wallet_address IS NOT NULL;