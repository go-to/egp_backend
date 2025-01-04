CREATE TABLE IF NOT EXISTS shops
(
    id                           INT AUTO_INCREMENT COMMENT 'ID'
        primary key,
    event_id                     INT                                 NOT NULL COMMENT 'イベントID',
    category_id                  INT                                 NOT NULL COMMENT 'カテゴリID',
    no                           INT                                 NOT NULL COMMENT 'No.',
    shop_name                    VARCHAR(255)                        NOT NULL COMMENT '店舗名',
    menu_name                    VARCHAR(255)                        NOT NULL COMMENT 'メニュー名',
    phone                        VARCHAR(255)                        NOT NULL COMMENT '電話番号',
    address                      VARCHAR(255)                        NOT NULL COMMENT '住所',
    business_days                VARCHAR(255)                        NOT NULL COMMENT '営業日',
    regular_holiday              VARCHAR(255)                        NOT NULL COMMENT '定休日',
    business_hours               VARCHAR(255)                        NOT NULL COMMENT '提供時間',
    charge_price                 VARCHAR(255)                        NOT NULL COMMENT 'チャージ',
    normalized_charge_price      INT                                 NULL COMMENT '(正規化済み)チャージ',
    single_price                 VARCHAR(255)                        NOT NULL COMMENT '逸品単品の値段',
    normalized_single_price      INT                                 NULL COMMENT '(正規化済み)逸品単品の値段',
    set_price                    VARCHAR(255)                        NOT NULL COMMENT 'セット料金',
    normalized_set_price         INT                                 NULL COMMENT '(正規化済み)セット料金',
    beer_type                    VARCHAR(255)                        NOT NULL COMMENT 'ビールの種類(銘柄)',
    needs_reservation            VARCHAR(255)                        NOT NULL COMMENT '予約の要・不要',
    normalized_needs_reservation TINYINT(1)                          NOT NULL COMMENT '(正規化済み)予約の要・不要',
    use_hachipay                 VARCHAR(255)                        NOT NULL COMMENT 'ハチペイ',
    normalized_use_hachipay      TINYINT(1)                          NOT NULL COMMENT '(正規化済み)ハチペイ',
    is_open_holiday              TINYINT(1)                          NOT NULL COMMENT '祝日営業フラグ',
    is_irregular_holiday         TINYINT(1)                          NOT NULL COMMENT '不定休フラグ',
    created_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at                   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_categories_id_fk
        FOREIGN KEY (category_id) REFERENCES categories (id),
    CONSTRAINT shops_events_id_fk
        FOREIGN KEY (event_id) REFERENCES events (id)
)
    COMMENT '店舗';
