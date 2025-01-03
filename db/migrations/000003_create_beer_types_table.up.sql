CREATE TABLE IF NOT EXISTS beer_types
(
    id         INT AUTO_INCREMENT COMMENT 'ID'
        PRIMARY KEY,
    name       VARCHAR(255)                        NOT NULL COMMENT '銘柄名',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL
)
    COMMENT 'ビールの種類';
