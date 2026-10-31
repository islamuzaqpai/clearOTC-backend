CREATE TABLE clearotc.deals
(
    id               UUID PRIMARY KEY,
    creator_id       UUID           NOT NULL,
    CONSTRAINT fk_deals_creator
        FOREIGN KEY (creator_id)
            REFERENCES clearotc.users (id),

    counterparty_id  UUID           NULL,
    CONSTRAINT fk_deals_counterparty
        FOREIGN KEY (counterparty_id)
            REFERENCES clearotc.users (id),

    contract_deal_id NUMERIC(78, 0) NULL,
    token_address    VARCHAR(42)    NULL,
    token_amount     NUMERIC(78, 0) NOT NULL,
    status           VARCHAR(30)    NOT NULL
        CHECK (status IN (
                          'PENDING',     ---Сделка создана---
                          'ACCEPTED',    ---Второй участник принял сделку---
                          'FUNDS_LOCKED',---Токены заблокированы контрактом---
                          'COMPLETED',   ---Сделка завершена---

                          'CANCELLED',   ---	Сделка отменена---
                          'EXPIRED',     ---Истёк срок действия---
                          'REFUNDED'     ---Токены возвращены---
            )),
    expires_at       TIMESTAMPTZ    NOT NULL,
    created_at       TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ    NOT NULL DEFAULT now(),
    deleted_at       TIMESTAMPTZ    NULL
);