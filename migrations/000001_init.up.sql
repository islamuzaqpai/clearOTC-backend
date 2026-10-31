CREATE SCHEMA IF NOT EXISTS clearotc;





-- CREATE TABLE clearotc.payments
-- (
--     id                 UUID PRIMARY KEY,
--     deal_id            UUID           NOT NULL,
--     CONSTRAINT fk_payments_deal
--         FOREIGN KEY (deal_id)
--             REFERENCES clearotc.deals (id),
--
--     provider           VARCHAR(50)    NOT NULL,
--     provider_reference VARCHAR(255)   NULL,
--     amount             NUMERIC(20, 2) NOT NULL,
--     currency           CHAR(3)        NOT NULL,
--     status             VARCHAR(30)    NOT NULL
--         CHECK ( status IN (
--             ''
--             ) ),
--
--     idempotency_key VARCHAR(255),
--     confirmed_at timestamptz NULL,
--     created_at       TIMESTAMPTZ    NOT NULL DEFAULT now(),
--     updated_at       TIMESTAMPTZ    NOT NULL DEFAULT now(),
--     deleted_at       TIMESTAMPTZ    NULL
-- )

--- SECOND MVP---


