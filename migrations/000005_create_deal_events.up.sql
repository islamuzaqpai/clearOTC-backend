CREATE TABLE clearotc.deal_events
(
    id            BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    deal_id       UUID        NOT NULL,
    CONSTRAINT fk_deal_events_deal
        FOREIGN KEY (deal_id)
            REFERENCES clearotc.deals (id),

    event_type    VARCHAR(50) NOT NULL,
    actor_user_id UUID        NULL,
    CONSTRAINT fk_deal_events_actor_user
        FOREIGN KEY (actor_user_id)
            REFERENCES clearotc.users (id),

    source        VARCHAR(30) NOT NULL
        CHECK (source IN (
                          'API',
                          'PAYMENT_PROVIDER',
                          'BLOCKCHAIN',
                          'SYSTEM'
            )),

    metadata      jsonb       NOT NULL DEFAULT '{}',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_deal_events_deal_created
    ON clearotc.deal_events (deal_id, created_at DESC);