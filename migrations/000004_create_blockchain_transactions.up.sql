CREATE TABLE clearotc.blockchain_transactions
(
    id            UUID PRIMARY KEY,
    deal_id       UUID               NOT NULL,
    CONSTRAINT fk_blockchain_transactions_deal
        FOREIGN KEY (deal_id)
            REFERENCES clearotc.deals (id),

    tx_hash       VARCHAR(66) UNIQUE NULL,
    status        VARCHAR(30)        NOT NULL DEFAULT 'PENDING'
        CHECK (status IN (
                          'PENDING',
                          'CONFIRMED',
                          'FAILED'
            )),

    block_number  BIGINT             NULL,
    tx_type       VARCHAR(30)        NOT NULL,
    error_message TEXT               NULL,
    created_at    TIMESTAMPTZ        NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ        NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMPTZ        NULL
);

CREATE INDEX idx_deals_creator_id
    ON clearotc.deals (creator_id);

CREATE INDEX idx_deals_counterparty_id
    ON clearotc.deals (counterparty_id);

CREATE INDEX idx_blockchain_transactions_deal_id
    ON clearotc.blockchain_transactions (deal_id);
