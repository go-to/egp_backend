CREATE TABLE IF NOT EXISTS egp.config
(
    id         SERIAL PRIMARY KEY,
    conf_name  VARCHAR(255)                        NOT NULL,
    conf_value VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE egp.config IS '設定値';
COMMENT ON COLUMN egp.config.id IS 'ID';
COMMENT ON COLUMN egp.config.conf_name IS '設定項目名';
COMMENT ON COLUMN egp.config.conf_value IS '設定値';
