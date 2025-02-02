CREATE TABLE IF NOT EXISTS egp.shops_beer_types_link
(
    id           SERIAL PRIMARY KEY,
    shop_id      INT                                 NOT NULL,
    beer_type_id INT                                 NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    CONSTRAINT shops_beer_types_link_beer_types_id_fk
        FOREIGN KEY (beer_type_id) REFERENCES egp.beer_types (id),
    CONSTRAINT shops_beer_types_link_shops_id_fk
        FOREIGN KEY (shop_id) REFERENCES egp.shops (id)
);
COMMENT ON TABLE egp.shops_beer_types_link IS '店舗-ビールの種類紐付け';
COMMENT ON COLUMN egp.shops_beer_types_link.id IS 'ID';
COMMENT ON COLUMN egp.shops_beer_types_link.shop_id IS '店舗ID';
COMMENT ON COLUMN egp.shops_beer_types_link.beer_type_id IS '銘柄ID';

CREATE INDEX shops_beer_types_link_beer_type_id_shop_id_index
    ON egp.shops_beer_types_link (beer_type_id, shop_id);
