CREATE TABLE IF NOT EXISTS egp.stamps_detail
(
    id         SERIAL
        PRIMARY KEY,
    user_id    VARCHAR(255) NOT NULL,
    shop_id    INTEGER      NOT NULL,
    stamped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT stamps_detail_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES egp.shops (id)
);
COMMENT ON TABLE egp.stamps_detail IS 'スタンプ詳細';
COMMENT ON COLUMN egp.stamps_detail.id IS 'ID';
COMMENT ON COLUMN egp.stamps_detail.user_id IS 'ユーザーID';
COMMENT ON COLUMN egp.stamps_detail.shop_id IS '店舗ID';
COMMENT ON COLUMN egp.stamps_detail.stamped_at IS 'スタンプ獲得日時';
