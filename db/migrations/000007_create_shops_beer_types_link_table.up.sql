CREATE TABLE IF NOT EXISTS shops_beer_types_link
(
    id           INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    shop_id      INT                                 NOT NULL COMMENT '店舗ID',
    beer_type_id INT                                 NOT NULL COMMENT '銘柄ID',
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_beer_types_link_beer_types_id_fk
        FOREIGN KEY (beer_type_id) REFERENCES beer_types (id),
    CONSTRAINT shops_beer_types_link_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES shops (id)
)
    COMMENT '店舗-ビールの種類紐付け';

CREATE INDEX shops_beer_types_link_beer_type_id_shop_id_index
    ON shops_beer_types_link (beer_type_id, shop_id);
