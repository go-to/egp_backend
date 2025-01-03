CREATE TABLE IF NOT EXISTS categories
(
    id         INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    name       VARCHAR(255)                        NOT NULL COMMENT 'カテゴリ名',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
)
    COMMENT 'カテゴリ';

