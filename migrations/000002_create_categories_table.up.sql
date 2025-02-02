CREATE TABLE IF NOT EXISTS egp.categories
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
);
COMMENT ON TABLE egp.categories IS 'カテゴリ';
COMMENT ON COLUMN egp.categories.id IS 'ID';
COMMENT ON COLUMN egp.categories.name IS 'カテゴリ名';
