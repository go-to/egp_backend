CREATE TABLE IF NOT EXISTS egp.stamps
(
    id              SERIAL
        PRIMARY KEY,
    user_id         VARCHAR(255) NOT NULL,
    shop_id         INTEGER      NOT NULL,
    number_of_times INTEGER      NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT stamps_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES egp.shops (id)
);
COMMENT ON TABLE egp.stamps IS 'スタンプ';
COMMENT ON COLUMN egp.stamps.id IS 'ID';
COMMENT ON COLUMN egp.stamps.user_id IS 'ユーザーID';
COMMENT ON COLUMN egp.stamps.shop_id IS '店舗ID';
COMMENT ON COLUMN egp.stamps.number_of_times IS '回数';
