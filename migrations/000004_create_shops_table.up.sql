CREATE TABLE IF NOT EXISTS egp.shops
(
    id                           SERIAL PRIMARY KEY,
    event_id                     INT                                 NOT NULL,
    category_id                  INT                                 NOT NULL,
    no                           INT                                 NOT NULL,
    shop_name                    VARCHAR(255)                        NOT NULL,
    menu_name                    VARCHAR(255)                        NOT NULL,
    menu_image_url               VARCHAR(255)                        NOT NULL,
    phone                        VARCHAR(255)                        NOT NULL,
    address                      VARCHAR(255)                        NOT NULL,
    business_days                VARCHAR(255)                        NOT NULL,
    regular_holiday              VARCHAR(255)                        NOT NULL,
    business_hours               VARCHAR(255)                        NOT NULL,
    charge_price                 VARCHAR(255)                        NOT NULL,
    normalized_charge_price      INT                                 NULL,
    single_price                 VARCHAR(255)                        NOT NULL,
    normalized_single_price      INT                                 NULL,
    set_price                    VARCHAR(255)                        NOT NULL,
    normalized_set_price         INT                                 NULL,
    beer_type                    VARCHAR(255)                        NOT NULL,
    needs_reservation            VARCHAR(255)                        NOT NULL,
    normalized_needs_reservation BOOLEAN                             NOT NULL,
    use_hachipay                 VARCHAR(255)                        NOT NULL,
    normalized_use_hachipay      BOOLEAN                             NOT NULL,
    is_open_holiday              BOOLEAN                             NOT NULL,
    is_irregular_holiday         BOOLEAN                             NOT NULL,
    created_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_categories_id_fk
        FOREIGN KEY (category_id) REFERENCES egp.categories (id),
    CONSTRAINT shops_events_id_fk
        FOREIGN KEY (event_id) REFERENCES egp.events (id)
);
COMMENT ON TABLE egp.shops IS '店舗';
COMMENT ON COLUMN egp.shops.id IS 'ID';
COMMENT ON COLUMN egp.shops.event_id IS 'イベントID';
COMMENT ON COLUMN egp.shops.category_id IS 'カテゴリID';
COMMENT ON COLUMN egp.shops.no IS 'No.';
COMMENT ON COLUMN egp.shops.shop_name IS '店舗名';
COMMENT ON COLUMN egp.shops.menu_name IS 'メニュー名';
COMMENT ON COLUMN egp.shops.menu_image_url IS 'メニュー画像URL';
COMMENT ON COLUMN egp.shops.phone IS '電話番号';
COMMENT ON COLUMN egp.shops.address IS '住所';
COMMENT ON COLUMN egp.shops.business_days IS '営業日';
COMMENT ON COLUMN egp.shops.regular_holiday IS '定休日';
COMMENT ON COLUMN egp.shops.business_hours IS '提供時間';
COMMENT ON COLUMN egp.shops.charge_price IS 'チャージ';
COMMENT ON COLUMN egp.shops.normalized_charge_price IS '(正規化済み)チャージ';
COMMENT ON COLUMN egp.shops.single_price IS '逸品単品の値段';
COMMENT ON COLUMN egp.shops.normalized_single_price IS '(正規化済み)逸品単品の値段';
COMMENT ON COLUMN egp.shops.set_price IS 'セット料金';
COMMENT ON COLUMN egp.shops.normalized_set_price IS '(正規化済み)セット料金';
COMMENT ON COLUMN egp.shops.beer_type IS 'ビールの種類(銘柄)';
COMMENT ON COLUMN egp.shops.needs_reservation IS '予約の要・不要';
COMMENT ON COLUMN egp.shops.normalized_needs_reservation IS '(正規化済み)予約の要・不要';
COMMENT ON COLUMN egp.shops.use_hachipay IS 'ハチペイ';
COMMENT ON COLUMN egp.shops.normalized_use_hachipay IS '(正規化済み)ハチペイ';
COMMENT ON COLUMN egp.shops.is_open_holiday IS '祝日営業フラグ';
COMMENT ON COLUMN egp.shops.is_irregular_holiday IS '不定休フラグ';
