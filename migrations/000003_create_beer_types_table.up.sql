CREATE TABLE IF NOT EXISTS egp.beer_types
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE egp.beer_types IS 'ビールの種類';
COMMENT ON COLUMN egp.beer_types.id IS 'ID';
COMMENT ON COLUMN egp.beer_types.name IS '銘柄名';