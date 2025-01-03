CREATE TABLE IF NOT EXISTS shops_location
(
    id         INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    shop_id    INT                                 NULL COMMENT '店舗ID',
    latitude   DOUBLE                              NOT NULL COMMENT '緯度',
    longitude  DOUBLE                              NOT NULL COMMENT '経度',
    location   POINT                               NOT NULL COMMENT '位置情報',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_location_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id)
)
    COMMENT '店舗位置情報';
