CREATE TABLE IF NOT EXISTS shops_time
(
    id          INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    shop_id     INT                                 NOT NULL COMMENT '店舗ID',
    week_number TINYINT(1)                          NOT NULL COMMENT '週番号',
    day_of_week TINYINT(1)                          NOT NULL COMMENT '曜日',
    start_time  TIME                                NOT NULL COMMENT '営業開始時間',
    end_time    TIME                                NOT NULL COMMENT '営業終了時間',
    is_holiday  TINYINT(1)                          NOT NULL COMMENT '定休日フラグ',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_time_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id)
)
    COMMENT '店舗営業時間';
